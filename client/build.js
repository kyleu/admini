const esbuild = require('esbuild');

esbuild.build({
  entryPoints: ['src/client.ts'],
  bundle: true,
  sourcemap: true,
  outfile: '../web/assets/client.js',
  watch: process.argv[2] === "watch" ? {
    onRebuild(error, result) {
      if (error) console.error('watch build failed:', error)
      else console.log('watch build succeeded:', result)
    }
  } : null,
  logLevel: "info"
}).catch((e) => console.error(e.message))
