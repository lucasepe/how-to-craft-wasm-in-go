'use strict';

function loadAndInitWA(waURL) {
    if (!WebAssembly.instantiateStreaming) { // polyfill
        WebAssembly.instantiateStreaming = async (resp, importObject) => {
            const source = await (await resp).arrayBuffer();
            return await WebAssembly.instantiate(source, importObject);
        };
    }
    const go = new Go();
    WebAssembly.instantiateStreaming(fetch(waURL), go.importObject)
    .then(async (result) => {
        await go.run(result.instance)
    });
}

