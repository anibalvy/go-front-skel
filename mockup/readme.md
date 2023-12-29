
# init (new project)
npx tailwindcss init
npx tailwindcss init tailwindcss-config.js    // if custom name is given, must be given as argument when compile
npx tailwindcss init --esm   // for ESM
npx tailwindcss init --ts    // for TypeScript
npx tailwindcss init -p      // to also generate a basic postcss.config.js file
npx tailwindcss init --full  // entire default configuration

# build
npx tailwindcss -c tailwind.config.js  -i css/style.css -o ./dist/output.css --watch


# build for production (if using mockup)
npx tailwindcss -c mockup/tailwind.config.js  -i mockup/css/style.css -o ./static/css/output.css --minify
