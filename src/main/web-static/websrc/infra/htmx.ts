function setHtmxRequestModifier() {
    document.addEventListener("htmx:configRequest", (event: CustomEvent) => {
        console.log("htmx:configRequest", event);
        const sourceElement = event.target;
        // TODO config a "param-sources" attibute on the sourceElement (one value is "path-context")
        // use it get param values and set same in the request
        // maybe get a list of possible params used in the url to be mapped
        // e.g. "portfolio" would scan the browser url for "/portfolio/1234" and set the value "1234" in the request
    });
}

export const htmxInfra = {
    init() {
        setHtmxRequestModifier();
    },
};