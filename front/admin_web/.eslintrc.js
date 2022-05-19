module.exports = {
    //解决错误： Parsing error: Unexpected token <
    root: true,
    parserOptions: {
        sourceType: 'module'
    },
    parser: "vue-eslint-parser",
    env: {
        browser: true,
        node: true,
        es6: true,
    },
    rules: {
        'no-console': 'off',
    },

    //全局变量
    globals: {
        AMap: true
    }
}