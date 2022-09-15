// .eslintrc.js
module.exports = {
    extends: ["plugin:vue/recommended", "plugin:prettier-vue/recommended"],

    settings: {
        "prettier-vue": {
            SFCBlocks: {
                template: true,

                script: true,
                style: true,

                customBlocks: {
                    docs: { lang: "markdown" },

                    config: { lang: "json" },

                    module: { lang: "js" },

                    comments: false,
                },
            },

            // Use prettierrc for prettier options or not (default: `true`)
            usePrettierrc: true,

            // Set the options for `prettier.getFileInfo`.
            // @see https://prettier.io/docs/en/api.html#prettiergetfileinfofilepath-options
            fileInfoOptions: {
                // Path to ignore file (default: `'.prettierignore'`)
                // Notice that the ignore file is only used for this plugin
                ignorePath: ".testignore",

                // Process the files in `node_modules` or not (default: `false`)
                withNodeModules: false,
            },
        },
    },

    rules: {
        "prettier-vue/prettier": [
            "error",
            {
                // Override all options of `prettier` here
                // @see https://prettier.io/docs/en/options.html
                printWidth: 100,
                singleQuote: true,
                semi: false,
                trailingComma: "es5",
            },
        ],
    },
};
