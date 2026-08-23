// eslint-disable-next-line @typescript-eslint/no-require-imports,no-undef
const { createProxyMiddleware } = require("http-proxy-middleware");
// eslint-disable-next-line @typescript-eslint/no-require-imports,no-undef
const fs = require("fs");
// eslint-disable-next-line @typescript-eslint/no-require-imports,no-undef
const path = require("path");

// Uses the host backend by default and accepts a Compose service URL when configured.
// Co-authored by: OpenCode and Igor Benicio de Mesquita
// eslint-disable-next-line no-undef
const apiProxyTarget = process.env.API_PROXY_TARGET || "http://localhost:8080/";

// Configures Parcel development middleware. For example, API_PROXY_TARGET=http://backend:8080/
// routes browser API requests through the Compose backend service.
// Co-authored by: OpenCode and Igor Benicio de Mesquita
// eslint-disable-next-line no-undef
module.exports = function (app) {

    // Proxy API requests to the backend
    app.use(
        createProxyMiddleware({ pathFilter: "/api", target: apiProxyTarget, changeOrigin: true }),
    );

    // Override default Parcel behavior to serve static files from the dist folder
    // return 404 error instead or redirecting to source
    app.use("/", (req, res, next) => {

        // eslint-disable-next-line no-undef
        const dir = path.join(__dirname, "dist");
        const requestURL = req.originalUrl;
        const filename = requestURL.split("?")[0];
        const isHTMXRequest = req.headers["hx-request"];

        if(isHTMXRequest && !fs.existsSync(path.join(dir, filename))) {
            res.statusCode = 404;
            next(`Resource not found: ${requestURL}`);
        } else {
            next();
        }
    });
};
