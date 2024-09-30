const resolve = require('@rollup/plugin-node-resolve');
const typescript = require('@rollup/plugin-typescript');
const commonjs = require('@rollup/plugin-commonjs');
const {terser} = require('@rollup/plugin-terser')
const copy = require("rollup-plugin-copy")
const serve = require('rollup-plugin-serve')

module.exports = [
    {
        input: './src/index.ts',
        output: [
            {
                dir: 'lib',
                format: 'cjs',
                entryFileNames: '[name].cjs.js',
                sourcemap: false, // 是否输出sourcemap
            },
            {
                dir: 'lib',
                format: 'esm',
                entryFileNames: '[name].esm.js',
                sourcemap: false, // 是否输出sourcemap
            },
            {
                dir: 'lib',
                format: 'umd',
                entryFileNames: '[name].umd.js',
                name: 'dgm',
                sourcemap: false,
                plugins: [terser],
            },
        ],
        plugins: [
            resolve(),
            commonjs(),
            typescript({module: "ESNext"}),
            copy({
                targets: [{src: 'README.md', dest: 'lib'}]
            }),
            serve({
                contentBase: ['public', 'lib'],
                port: 8081
            })
        ],
    }
]

