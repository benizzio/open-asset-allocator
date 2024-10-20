type Level =
    "debug" |
    "error" |
    "info" |
    "log" |
    "trace" |
    "warn";

type Levels = {
    DEBUG: Level,
    ERROR: Level,
    INFO: Level,
    LOG: Level,
    TRACE: Level,
    WARN: Level
};

const LogLevel: Levels = {
    DEBUG: "debug",
    ERROR: "error",
    INFO: "info",
    LOG: "log",
    TRACE: "trace",
    WARN: "warn",
};

const logger = (logLevel: Level, ...args: unknown[]) => {
    console[logLevel](...args);
};

export { logger, LogLevel };