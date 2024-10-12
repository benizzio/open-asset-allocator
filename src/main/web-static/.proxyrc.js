// eslint-disable-next-line @typescript-eslint/no-require-imports,no-undef
const { createProxyMiddleware } = require("http-proxy-middleware");
// eslint-disable-next-line @typescript-eslint/no-require-imports,no-undef
const fs = require("fs");
// eslint-disable-next-line @typescript-eslint/no-require-imports,no-undef
const path = require("path");

// eslint-disable-next-line no-undef
module.exports = function (app) {

    app.use(
        createProxyMiddleware("/api", {
            target: "http://localhost:8080/",
            pathRewrite: { "^/api": "" },
        }),
    );

    app.use("/", (req, res, next) => {

        // eslint-disable-next-line no-undef
        const dir = path.join(__dirname, "dist");
        const requestURL = req.originalUrl;
        const filename = requestURL.split("?")[0];

        if (!fs.existsSync(path.join(dir, filename))) {
            res.statusCode = 404;
            next(`Resource not found: ${requestURL}`);
        }
        else {
            next();
        }
    });
};