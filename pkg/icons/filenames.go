package icons

var IconFileName = map[string]*IconInfo{
	".pug-lintrc":                         IconSet["pug"],
	".pug-lintrc.js":                      IconSet["pug"],
	".pug-lintrc.json":                    IconSet["pug"],
	".jscsrc":                             IconSet["json"],
	".jshintrc":                           IconSet["json"],
	"composer.lock":                       IconSet["json"],
	".jsbeautifyrc":                       IconSet["json"],
	".esformatter":                        IconSet["json"],
	"cdp.pid":                             IconSet["json"],
	".mjmlconfig":                         IconSet["json"],
	".htaccess":                           IconSet["xml"],
	".jshintignore":                       IconSet["settings"],
	".buildignore":                        IconSet["settings"],
	".mrconfig":                           IconSet["settings"],
	".yardopts":                           IconSet["settings"],
	"manifest.mf":                         IconSet["settings"],
	".clang-format":                       IconSet["settings"],
	".clang-tidy":                         IconSet["settings"],
	"go.mod":                              IconSet["go-mod"],
	"go.sum":                              IconSet["go-mod"],
	"requirements.txt":                    IconSet["python-misc"],
	"pipfile":                             IconSet["python-misc"],
	".python-version":                     IconSet["python-misc"],
	"manifest.in":                         IconSet["python-misc"],
	"gradle.properties":                   IconSet["gradle"],
	"gradlew":                             IconSet["gradle"],
	"gradle-wrapper.properties":           IconSet["gradle"],
	"license":                             IconSet["certificate"],
	"license.md":                          IconSet["certificate"],
	"license.txt":                         IconSet["certificate"],
	"licence":                             IconSet["certificate"],
	"licence.md":                          IconSet["certificate"],
	"licence.txt":                         IconSet["certificate"],
	"unlicense":                           IconSet["certificate"],
	"unlicense.md":                        IconSet["certificate"],
	"unlicense.txt":                       IconSet["certificate"],
	".htpasswd":                           IconSet["key"],
	"gemfile":                             IconSet["gemfile"],
	"dockerfile":                          IconSet["docker"],
	"dockerfile.prod":                     IconSet["docker"],
	"dockerfile.production":               IconSet["docker"],
	"docker-compose.yml":                  IconSet["docker"],
	"docker-compose.yaml":                 IconSet["docker"],
	"docker-compose.dev.yml":              IconSet["docker"],
	"docker-compose.local.yml":            IconSet["docker"],
	"docker-compose.ci.yml":               IconSet["docker"],
	"docker-compose.override.yml":         IconSet["docker"],
	"docker-compose.staging.yml":          IconSet["docker"],
	"docker-compose.prod.yml":             IconSet["docker"],
	"docker-compose.production.yml":       IconSet["docker"],
	"docker-compose.test.yml":             IconSet["docker"],
	".mailmap":                            IconSet["email"],
	".graphqlconfig":                      IconSet["graphql"],
	".gitignore":                          IconSet["git"],
	".gitconfig":                          IconSet["git"],
	".gitattributes":                      IconSet["git"],
	".gitmodules":                         IconSet["git"],
	".gitkeep":                            IconSet["git"],
	"git-history":                         IconSet["git"],
	".luacheckrc":                         IconSet["lua"],
	".Rhistory":                           IconSet["r"],
	"cmakelists.txt":                      IconSet["cmake"],
	"cmakecache.txt":                      IconSet["cmake"],
	"vue.config.js":                       IconSet["vue-config"],
	"vue.config.ts":                       IconSet["vue-config"],
	"nuxt.config.js":                      IconSet["nuxt"],
	"nuxt.config.ts":                      IconSet["nuxt"],
	"security.md":                         IconSet["lock"],
	"security.txt":                        IconSet["lock"],
	"security":                            IconSet["lock"],
	"vercel.json":                         IconSet["vercel"],
	".vercelignore":                       IconSet["vercel"],
	"now.json":                            IconSet["vercel"],
	".nowignore":                          IconSet["vercel"],
	"postcss.config.js":                   IconSet["postcss"],
	".postcssrc.js":                       IconSet["postcss"],
	".postcssrc":                          IconSet["postcss"],
	".postcssrc.json":                     IconSet["postcss"],
	".postcssrc.yml":                      IconSet["postcss"],
	"CNAME":                               IconSet["http"],
	"webpack.js":                          IconSet["webpack"],
	"webpack.ts":                          IconSet["webpack"],
	"webpack.base.js":                     IconSet["webpack"],
	"webpack.base.ts":                     IconSet["webpack"],
	"webpack.config.js":                   IconSet["webpack"],
	"webpack.config.ts":                   IconSet["webpack"],
	"webpack.common.js":                   IconSet["webpack"],
	"webpack.common.ts":                   IconSet["webpack"],
	"webpack.config.common.js":            IconSet["webpack"],
	"webpack.config.common.ts":            IconSet["webpack"],
	"webpack.config.common.babel.js":      IconSet["webpack"],
	"webpack.config.common.babel.ts":      IconSet["webpack"],
	"webpack.dev.js":                      IconSet["webpack"],
	"webpack.dev.ts":                      IconSet["webpack"],
	"webpack.development.js":              IconSet["webpack"],
	"webpack.development.ts":              IconSet["webpack"],
	"webpack.config.dev.js":               IconSet["webpack"],
	"webpack.config.dev.ts":               IconSet["webpack"],
	"webpack.config.dev.babel.js":         IconSet["webpack"],
	"webpack.config.dev.babel.ts":         IconSet["webpack"],
	"webpack.prod.js":                     IconSet["webpack"],
	"webpack.prod.ts":                     IconSet["webpack"],
	"webpack.production.js":               IconSet["webpack"],
	"webpack.production.ts":               IconSet["webpack"],
	"webpack.server.js":                   IconSet["webpack"],
	"webpack.server.ts":                   IconSet["webpack"],
	"webpack.client.js":                   IconSet["webpack"],
	"webpack.client.ts":                   IconSet["webpack"],
	"webpack.config.server.js":            IconSet["webpack"],
	"webpack.config.server.ts":            IconSet["webpack"],
	"webpack.config.client.js":            IconSet["webpack"],
	"webpack.config.client.ts":            IconSet["webpack"],
	"webpack.config.production.babel.js":  IconSet["webpack"],
	"webpack.config.production.babel.ts":  IconSet["webpack"],
	"webpack.config.prod.babel.js":        IconSet["webpack"],
	"webpack.config.prod.babel.ts":        IconSet["webpack"],
	"webpack.config.prod.js":              IconSet["webpack"],
	"webpack.config.prod.ts":              IconSet["webpack"],
	"webpack.config.production.js":        IconSet["webpack"],
	"webpack.config.production.ts":        IconSet["webpack"],
	"webpack.config.staging.js":           IconSet["webpack"],
	"webpack.config.staging.ts":           IconSet["webpack"],
	"webpack.config.babel.js":             IconSet["webpack"],
	"webpack.config.babel.ts":             IconSet["webpack"],
	"webpack.config.base.babel.js":        IconSet["webpack"],
	"webpack.config.base.babel.ts":        IconSet["webpack"],
	"webpack.config.base.js":              IconSet["webpack"],
	"webpack.config.base.ts":              IconSet["webpack"],
	"webpack.config.staging.babel.js":     IconSet["webpack"],
	"webpack.config.staging.babel.ts":     IconSet["webpack"],
	"webpack.config.coffee":               IconSet["webpack"],
	"webpack.config.test.js":              IconSet["webpack"],
	"webpack.config.test.ts":              IconSet["webpack"],
	"webpack.config.vendor.js":            IconSet["webpack"],
	"webpack.config.vendor.ts":            IconSet["webpack"],
	"webpack.config.vendor.production.js": IconSet["webpack"],
	"webpack.config.vendor.production.ts": IconSet["webpack"],
	"webpack.test.js":                     IconSet["webpack"],
	"webpack.test.ts":                     IconSet["webpack"],
	"webpack.dist.js":                     IconSet["webpack"],
	"webpack.dist.ts":                     IconSet["webpack"],
	"webpackfile.js":                      IconSet["webpack"],
	"webpackfile.ts":                      IconSet["webpack"],
	"ionic.config.json":                   IconSet["ionic"],
	".io-config.json":                     IconSet["ionic"],
	"gulpfile.js":                         IconSet["gulp"],
	"gulpfile.mjs":                        IconSet["gulp"],
	"gulpfile.ts":                         IconSet["gulp"],
	"gulpfile.babel.js":                   IconSet["gulp"],
	"package.json":                        IconSet["nodejs"],
	"package-lock.json":                   IconSet["nodejs"],
	".nvmrc":                              IconSet["nodejs"],
	".esmrc":                              IconSet["nodejs"],
	".node-version":                       IconSet["nodejs"],
	".npmignore":                          IconSet["npm"],
	".npmrc":                              IconSet["npm"],
	".yarnrc":                             IconSet["yarn"],
	"yarn.lock":                           IconSet["yarn"],
	".yarnclean":                          IconSet["yarn"],
	".yarn-integrity":                     IconSet["yarn"],
	"yarn-error.log":                      IconSet["yarn"],
	".yarnrc.yml":                         IconSet["yarn"],
	".yarnrc.yaml":                        IconSet["yarn"],
	"androidmanifest.xml":                 IconSet["android"],
	".env.defaults":                       IconSet["tune"],
	".env.example":                        IconSet["tune"],
	".env.sample":                         IconSet["tune"],
	".env.schema":                         IconSet["tune"],
	".env.local":                          IconSet["tune"],
	".env.dev":                            IconSet["tune"],
	".env.development":                    IconSet["tune"],
	".env.qa":                             IconSet["tune"],
	".env.prod":                           IconSet["tune"],
	".env.production":                     IconSet["tune"],
	".env.staging":                        IconSet["tune"],
	".env.preview":                        IconSet["tune"],
	".env.test":                           IconSet["tune"],
	".env.testing":                        IconSet["tune"],
	".env.development.local":              IconSet["tune"],
	".env.qa.local":                       IconSet["tune"],
	".env.production.local":               IconSet["tune"],
	".env.staging.local":                  IconSet["tune"],
	".env.test.local":                     IconSet["tune"],
	".babelrc":                            IconSet["babel"],
	".babelrc.js":                         IconSet["babel"],
	".babelrc.json":                       IconSet["babel"],
	"babel.config.json":                   IconSet["babel"],
	"babel.config.js":                     IconSet["babel"],
	"contributing.md":                     IconSet["contributing"],
	"readme.md":                           IconSet["readme"],
	"readme.txt":                          IconSet["readme"],
	"readme":                              IconSet["readme"],
	"changelog":                           IconSet["changelog"],
	"changelog.md":                        IconSet["changelog"],
	"changelog.txt":                       IconSet["changelog"],
	"changes":                             IconSet["changelog"],
	"changes.md":                          IconSet["changelog"],
	"changes.txt":                         IconSet["changelog"],
	"credits":                             IconSet["credits"],
	"credits.txt":                         IconSet["credits"],
	"credits.md":                          IconSet["credits"],
	"authors":                             IconSet["authors"],
	"authors.md":                          IconSet["authors"],
	"authors.txt":                         IconSet["authors"],
	"favicon.ico":                         IconSet["favicon"],
	"karma.conf.js":                       IconSet["karma"],
	"karma.conf.ts":                       IconSet["karma"],
	"karma.conf.coffee":                   IconSet["karma"],
	"karma.config.js":                     IconSet["karma"],
	"karma.config.ts":                     IconSet["karma"],
	"karma-main.js":                       IconSet["karma"],
	"karma-main.ts":                       IconSet["karma"],
	".travis.yml":                         IconSet["travis"],
	".codecov.yml":                        IconSet["codecov"],
	"codecov.yml":                         IconSet["codecov"],
	"protractor.conf.js":                  IconSet["protractor"],
	"protractor.conf.ts":                  IconSet["protractor"],
	"protractor.conf.coffee":              IconSet["protractor"],
	"protractor.config.js":                IconSet["protractor"],
	"protractor.config.ts":                IconSet["protractor"],
	"procfile":                            IconSet["heroku"],
	"procfile.windows":                    IconSet["heroku"],
	".bowerrc":                            IconSet["bower"],
	"bower.json":                          IconSet["bower"],
	".eslintrc.js":                        IconSet["eslint"],
	".eslintrc.cjs":                       IconSet["eslint"],
	".eslintrc.yaml":                      IconSet["eslint"],
	".eslintrc.yml":                       IconSet["eslint"],
	".eslintrc.json":                      IconSet["eslint"],
	".eslintrc":                           IconSet["eslint"],
	".eslintignore":                       IconSet["eslint"],
	".eslintcache":                        IconSet["eslint"],
	"code_of_conduct.md":                  IconSet["conduct"],
	"code_of_conduct.txt":                 IconSet["conduct"],
	"mocha.opts":                          IconSet["mocha"],
	".mocharc.yml":                        IconSet["mocha"],
	".mocharc.yaml":                       IconSet["mocha"],
	".mocharc.js":                         IconSet["mocha"],
	".mocharc.json":                       IconSet["mocha"],
	".mocharc.jsonc":                      IconSet["mocha"],
	"jenkinsfile":                         IconSet["jenkins"],
	"firebase.json":                       IconSet["firebase"],
	".firebaserc":                         IconSet["firebase"],
	"firestore.rules":                     IconSet["firebase"],
	"firestore.indexes.json":              IconSet["firebase"],
	".stylelintrc":                        IconSet["stylelint"],
	"stylelint.config.js":                 IconSet["stylelint"],
	".stylelintrc.json":                   IconSet["stylelint"],
	".stylelintrc.yaml":                   IconSet["stylelint"],
	".stylelintrc.yml":                    IconSet["stylelint"],
	".stylelintrc.js":                     IconSet["stylelint"],
	".stylelintignore":                    IconSet["stylelint"],
	".codeclimate.yml":                    IconSet["code-climate"],
	".prettierrc":                         IconSet["prettier"],
	"prettier.config.js":                  IconSet["prettier"],
	".prettierrc.js":                      IconSet["prettier"],
	".prettierrc.json":                    IconSet["prettier"],
	".prettierrc.yaml":                    IconSet["prettier"],
	".prettierrc.yml":                     IconSet["prettier"],
	".prettierignore":                     IconSet["prettier"],
	"gruntfile.js":                        IconSet["grunt"],
	"gruntfile.ts":                        IconSet["grunt"],
	"gruntfile.coffee":                    IconSet["grunt"],
	"gruntfile.babel.js":                  IconSet["grunt"],
	"gruntfile.babel.ts":                  IconSet["grunt"],
	"gruntfile.babel.coffee":              IconSet["grunt"],
	"jest.config.js":                      IconSet["jest"],
	"jest.config.ts":                      IconSet["jest"],
	"jest.config.cjs":                     IconSet["jest"],
	"jest.config.mjs":                     IconSet["jest"],
	"jest.config.json":                    IconSet["jest"],
	"jest.e2e.config.js":                  IconSet["jest"],
	"jest.e2e.config.ts":                  IconSet["jest"],
	"jest.e2e.config.cjs":                 IconSet["jest"],
	"jest.e2e.config.mjs":                 IconSet["jest"],
	"jest.e2e.config.json":                IconSet["jest"],
	"jest.setup.js":                       IconSet["jest"],
	"jest.setup.ts":                       IconSet["jest"],
	"jest.json":                           IconSet["jest"],
	".jestrc":                             IconSet["jest"],
	".jestrc.js":                          IconSet["jest"],
	".jestrc.json":                        IconSet["jest"],
	"jest.teardown.js":                    IconSet["jest"],
	"fastfile":                            IconSet["fastlane"],
	"appfile":                             IconSet["fastlane"],
	".helmignore":                         IconSet["helm"],
	"makefile":                            IconSet["makefile"],
	".releaserc":                          IconSet["semantic-release"],
	".releaserc.yaml":                     IconSet["semantic-release"],
	".releaserc.yml":                      IconSet["semantic-release"],
	".releaserc.json":                     IconSet["semantic-release"],
	".releaserc.js":                       IconSet["semantic-release"],
	"release.config.js":                   IconSet["semantic-release"],
	"bitbucket-pipelines.yaml":            IconSet["bitbucket"],
	"bitbucket-pipelines.yml":             IconSet["bitbucket"],
	"azure-pipelines.yml":                 IconSet["azure-pipelines"],
	"azure-pipelines.yaml":                IconSet["azure-pipelines"],
	"vagrantfile":                         IconSet["vagrant"],
	"tailwind.js":                         IconSet["tailwindcss"],
	"tailwind.config.js":                  IconSet["tailwindcss"],
	"codeowners":                          IconSet["codeowners"],
	".gcloudignore":                       IconSet["gcp"],
	".huskyrc":                            IconSet["husky"],
	"husky.config.js":                     IconSet["husky"],
	".huskyrc.json":                       IconSet["husky"],
	".huskyrc.js":                         IconSet["husky"],
	".huskyrc.yaml":                       IconSet["husky"],
	".huskyrc.yml":                        IconSet["husky"],
	".commitlintrc":                       IconSet["commitlint"],
	".commitlintrc.js":                    IconSet["commitlint"],
	"commitlint.config.js":                IconSet["commitlint"],
	".commitlintrc.json":                  IconSet["commitlint"],
	".commitlint.yaml":                    IconSet["commitlint"],
	".commitlint.yml":                     IconSet["commitlint"],
	"dune":                                IconSet["dune"],
	"dune-project":                        IconSet["dune"],
	"roadmap.md":                          IconSet["roadmap"],
	"roadmap.txt":                         IconSet["roadmap"],
	"timeline.md":                         IconSet["roadmap"],
	"timeline.txt":                        IconSet["roadmap"],
	"milestones.md":                       IconSet["roadmap"],
	"milestones.txt":                      IconSet["roadmap"],
	"nuget.config":                        IconSet["nuget"],
	".nuspec":                             IconSet["nuget"],
	"nuget.exe":                           IconSet["nuget"],
	"stryker.conf.js":                     IconSet["stryker"],
	"stryker.conf.json":                   IconSet["stryker"],
	".modernizrrc":                        IconSet["modernizr"],
	".modernizrrc.js":                     IconSet["modernizr"],
	".modernizrrc.json":                   IconSet["modernizr"],
	"routing.ts":                          IconSet["routing"],
	"routing.tsx":                         IconSet["routing"],
	"routing.js":                          IconSet["routing"],
	"routing.jsx":                         IconSet["routing"],
	"routes.ts":                           IconSet["routing"],
	"routes.tsx":                          IconSet["routing"],
	"routes.js":                           IconSet["routing"],
	"routes.jsx":                          IconSet["routing"],
	// ".vfl":                                IconSet["vfl"],
	// ".kl":                                 IconSet["kl"],
	// "project.graphcool":                   IconSet["graphcool"],
	// ".flowconfig":                         IconSet["flow"],
	// ".bithoundrc":                         IconSet["bithound"],
	// ".appveyor.yml":                       IconSet["appveyor"],
	// "appveyor.yml":                        IconSet["appveyor"],
	// "fuse.js":                             IconSet["fusebox"],
	// ".editorconfig":                       IconSet["editorconfig"],
	// ".watchmanconfig":                     IconSet["watchman"],
	// "aurelia.json":                        IconSet["aurelia"],
	// "rollup.config.js":                    IconSet["rollup"],
	// "rollup.config.ts":                    IconSet["rollup"],
	// "rollup-config.js":                    IconSet["rollup"],
	// "rollup-config.ts":                    IconSet["rollup"],
	// "rollup.config.common.js":             IconSet["rollup"],
	// "rollup.config.common.ts":             IconSet["rollup"],
	// "rollup.config.base.js":               IconSet["rollup"],
	// "rollup.config.base.ts":               IconSet["rollup"],
	// "rollup.config.prod.js":               IconSet["rollup"],
	// "rollup.config.prod.ts":               IconSet["rollup"],
	// "rollup.config.dev.js":                IconSet["rollup"],
	// "rollup.config.dev.ts":                IconSet["rollup"],
	// "rollup.config.prod.vendor.js":        IconSet["rollup"],
	// "rollup.config.prod.vendor.ts":        IconSet["rollup"],
	// ".hhconfig":                           IconSet["hack"],
	// "apollo.config.js":                    IconSet["apollo"],
	// "nodemon.json":                        IconSet["nodemon"],
	// "nodemon-debug.json":                  IconSet["nodemon"],
	// ".hintrc":                             IconSet["webhint"],
	// "browserslist":                        IconSet["browserlist"],
	// ".browserslistrc":                     IconSet["browserlist"],
	// ".snyk":                               IconSet["snyk"],
	// ".drone.yml":                          IconSet["drone"],
	// ".sequelizerc":                        IconSet["sequelize"],
	// "gatsby.config.js":                    IconSet["gatsby"],
	// "gatsby-config.js":                    IconSet["gatsby"],
	// "gatsby-node.js":                      IconSet["gatsby"],
	// "gatsby-browser.js":                   IconSet["gatsby"],
	// "gatsby-ssr.js":                       IconSet["gatsby"],
	// ".wakatime-project":                   IconSet["wakatime"],
	// "circle.yml":                          IconSet["circleci"],
	// ".cfignore":                           IconSet["cloudfoundry"],
	// "wallaby.js":                          IconSet["wallaby"],
	// "wallaby.conf.js":                     IconSet["wallaby"],
	// "stencil.config.js":                   IconSet["stencil"],
	// "stencil.config.ts":                   IconSet["stencil"],
	// ".bazelignore":                        IconSet["bazel"],
	// ".bazelrc":                            IconSet["bazel"],
	// "prisma.yml":                          IconSet["prisma"],
	// ".nycrc":                              IconSet["istanbul"],
	// ".nycrc.json":                         IconSet["istanbul"],
	// "buildkite.yml":                       IconSet["buildkite"],
	// "buildkite.yaml":                      IconSet["buildkite"],
	// "netlify.json":                        IconSet["netlify"],
	// "netlify.yml":                         IconSet["netlify"],
	// "netlify.yaml":                        IconSet["netlify"],
	// "netlify.toml":                        IconSet["netlify"],
	// "nest-cli.json":                       IconSet["nest"],
	// ".nest-cli.json":                      IconSet["nest"],
	// "nestconfig.json":                     IconSet["nest"],
	// ".nestconfig.json":                    IconSet["nest"],
	// ".percy.yml":                          IconSet["percy"],
	// ".gitpod.yml":                         IconSet["gitpod"],
	// "tiltfile":                            IconSet["tilt"],
	// "capacitor.config.json":               IconSet["capacitor"],
	// ".adonisrc.json":                      IconSet["adonis"],
	// "ace":                                 IconSet["adonis"],
	// "meson.build":                         IconSet["meson"],
	// ".buckconfig":                         IconSet["buck"],
	// "nx.json":                             IconSet["nrwl"],
	// ".slugignore":                         IconSet["slug"],
}
