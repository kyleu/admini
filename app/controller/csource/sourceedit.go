package csource

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
	"path/filepath"
	"strconv"

	"github.com/pkg/errors"
	"github.com/valyala/fasthttp"

	"admini.dev/admini/app"
	"admini.dev/admini/app/controller"
	"admini.dev/admini/app/controller/cutil"
	"admini.dev/admini/app/lib/database"
	"admini.dev/admini/app/lib/filesystem"
	"admini.dev/admini/app/lib/schema"
	"admini.dev/admini/app/source"
	"admini.dev/admini/app/util"
	"admini.dev/admini/assets"
	"admini.dev/admini/views/vsource"
)

func SourceNew(rc *fasthttp.RequestCtx) {
	controller.Act("source.new", rc, func(as *app.State, ps *cutil.PageState) (string, error) {
		ps.Title = "New Source"
		t := schema.OriginPostgres
		s := &source.Source{Type: t}
		ps.Data = s
		return controller.Render(rc, as, &vsource.New{Origin: t}, ps, "sources", "New")
	})
}

func SourceExample(rc *fasthttp.RequestCtx) {
	controller.Act("source.example", rc, func(as *app.State, ps *cutil.PageState) (string, error) {
		ps.Title = "Example Database"

		ent, err := assets.Embed("example.sqlite.gz")
		if err != nil {
			return "", errors.Wrap(err, "unable to load embedded example database")
		}
		b := bytes.NewBuffer(ent.Bytes)
		zr, err := gzip.NewReader(b)
		if err != nil {
			return "", err
		}
		defer func() { _ = zr.Close() }()
		out, err := io.ReadAll(zr)
		if err != nil {
			return "", errors.Wrap(err, "unable to decompress embedded example database")
		}

		fn := filepath.Join(as.Files.Root(), "example.sqlite")
		fpath := filepath.Join(as.Files.Root(), "example.sqlite")

		err = as.Files.WriteFile(fn, out, filesystem.DefaultMode, true)
		if err != nil {
			return "", errors.Wrapf(err, "unable to write example database to [%s]", fpath)
		}

		ret := as.Services.Sources.NewSource("example", "Example", "star", "Example music database", schema.OriginSQLite)
		ret.Config = []byte(util.ToJSON(map[string]any{"file": fpath}))
		err = as.Services.Sources.Save(ret, false, ps.Logger)
		if err != nil {
			return "", errors.Wrap(err, "unable to save example database")
		}
		ps.Data = ret
		return controller.FlashAndRedir(true, "saved example source", "/source/example", rc, ps)
	})
}

func SourceInsert(rc *fasthttp.RequestCtx) {
	controller.Act("source.insert", rc, func(as *app.State, ps *cutil.PageState) (string, error) {
		frm, err := cutil.ParseForm(rc)
		if err != nil {
			return "", errors.Wrap(err, "unable to parse form")
		}

		key, err := frm.GetString("key", false)
		if err != nil {
			return controller.FlashAndRedir(false, err.Error(), "/source/_new", rc, ps)
		}
		title := frm.GetStringOpt("title")
		icon := frm.GetStringOpt("icon")
		description := frm.GetStringOpt("description")
		typ := frm.GetStringOpt("type")
		ret := as.Services.Sources.NewSource(key, title, icon, description, schema.OriginFromString(typ))
		err = as.Services.Sources.Save(ret, false, ps.Logger)
		if err != nil {
			return "", errors.Wrap(err, "unable to save source")
		}
		return controller.FlashAndRedir(true, "saved new source", fmt.Sprintf("/source/%s", key), rc, ps)
	})
}

func SourceEdit(rc *fasthttp.RequestCtx) {
	controller.Act("source.edit", rc, func(as *app.State, ps *cutil.PageState) (string, error) {
		key, err := cutil.RCRequiredString(rc, "key", false)
		if err != nil {
			return "", err
		}
		src, err := as.Services.Sources.Load(key, false, ps.Logger)
		if err != nil {
			return "", errors.Wrapf(err, "unable to load source [%s]", key)
		}
		ps.Title = fmt.Sprintf("Edit [%s]", src.Name())
		ps.Data = src

		switch src.Type {
		case schema.OriginPostgres, schema.OriginSQLite, schema.OriginMySQL, schema.OriginSQLServer:
			page := &vsource.Edit{Source: src}
			return controller.Render(rc, as, page, ps, "sources", src.Key, "Edit")
		default:
			msg := fmt.Sprintf("unhandled source type [%s]", src.Type.String())
			return controller.FlashAndRedir(false, msg, "/source", rc, ps)
		}
	})
}

func SourceSave(rc *fasthttp.RequestCtx) {
	controller.Act("source.save", rc, func(as *app.State, ps *cutil.PageState) (string, error) {
		frm, err := cutil.ParseForm(rc)
		if err != nil {
			return "", errors.Wrap(err, "unable to parse form")
		}

		key, err := cutil.RCRequiredString(rc, "key", false)
		if err != nil {
			return "", err
		}
		src, err := as.Services.Sources.Load(key, false, ps.Logger)
		if err != nil {
			return "", errors.Wrapf(err, "unable to load source [%s]", key)
		}

		src.Title = frm.GetStringOpt("title")
		src.Icon = frm.GetStringOpt("icon")
		src.Description = frm.GetStringOpt("description")

		switch src.Type {
		case schema.OriginMySQL:
			ps := frm.GetStringOpt("port")
			params := &database.MySQLParams{}
			params.Host = frm.GetStringOpt("host")
			if ps != "" {
				params.Port, _ = strconv.Atoi(ps)
			}
			params.Username = frm.GetStringOpt("username")
			params.Password = frm.GetStringOpt("password")
			params.Database = frm.GetStringOpt("database")
			params.Schema = frm.GetStringOpt("schema")
			params.Debug, _ = frm.GetBool("debug", true)

			src.Config = util.ToJSONBytes(params, true)
		case schema.OriginPostgres:
			params := &database.PostgresParams{}
			params.Host = frm.GetStringOpt("host")
			if ps := frm.GetStringOpt("port"); ps != "" {
				params.Port, _ = strconv.Atoi(ps)
			}
			params.Username = frm.GetStringOpt("username")
			params.Password = frm.GetStringOpt("password")
			params.Database = frm.GetStringOpt("database")
			params.Schema = frm.GetStringOpt("schema")
			params.Debug, _ = frm.GetBool("debug", true)

			src.Config = util.ToJSONBytes(params, true)
		case schema.OriginSQLite:
			params := &database.SQLiteParams{}
			params.File = frm.GetStringOpt("file")
			params.Schema = frm.GetStringOpt("schema")
			params.Debug, _ = frm.GetBool("debug", true)

			src.Config = util.ToJSONBytes(params, true)
		case schema.OriginSQLServer:
			params := &database.SQLServerParams{}
			params.Host = frm.GetStringOpt("host")
			if ps := frm.GetStringOpt("port"); ps != "" {
				params.Port, _ = strconv.Atoi(ps)
			}
			params.Username = frm.GetStringOpt("username")
			params.Password = frm.GetStringOpt("password")
			params.Database = frm.GetStringOpt("database")
			params.Schema = frm.GetStringOpt("schema")
			params.Debug, _ = frm.GetBool("debug", true)

			src.Config = util.ToJSONBytes(params, true)
		default:
			return "", errors.Errorf("unable to parse config for source type [%s]", src.Type.String())
		}

		err = as.Services.Sources.Save(src, true, ps.Logger)
		if err != nil {
			return "", errors.Wrapf(err, "unable to save source [%s]", key)
		}

		msg := fmt.Sprintf(`saved source %q`, key)
		return controller.FlashAndRedir(true, msg, fmt.Sprintf("/source/%s", key), rc, ps)
	})
}

func SourceDelete(rc *fasthttp.RequestCtx) {
	controller.Act("source.delete", rc, func(as *app.State, ps *cutil.PageState) (string, error) {
		key, err := cutil.RCRequiredString(rc, "key", false)
		if err != nil {
			return "", err
		}
		err = as.Services.Sources.Delete(key, ps.Logger)
		if err != nil {
			return "", errors.Wrapf(err, "unable to delete source [%s]", key)
		}

		msg := fmt.Sprintf(`deleted source %q`, key)
		return controller.FlashAndRedir(true, msg, "/source", rc, ps)
	})
}
