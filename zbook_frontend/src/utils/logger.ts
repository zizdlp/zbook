import chalk from "chalk";

class Logger {
  private log(
    level: "info" | "warn" | "error",
    message: string,
    status?: number
  ) {
    const timestamp = new Date().toISOString();
    const errorType = chalk.bold(level.toUpperCase());
    const statusMessage = status ? ` [Status: ${status}]` : "";

    let logMessage = `[${timestamp}] [${errorType}]${statusMessage} ${message}`;

    switch (level) {
      case "info":
        logMessage = logMessage.replace(errorType, chalk.blue(errorType));
        if (status)
          logMessage = logMessage.replace(
            `[Status: ${status}]`,
            chalk.blue(`[Status: ${status}]`)
          );
        break;
      case "warn":
        logMessage = logMessage.replace(errorType, chalk.yellow(errorType));
        if (status)
          logMessage = logMessage.replace(
            `[Status: ${status}]`,
            chalk.yellow(`[Status: ${status}]`)
          );
        break;
      case "error":
        logMessage = logMessage.replace(errorType, chalk.red(errorType));
        if (status)
          logMessage = logMessage.replace(
            `[Status: ${status}]`,
            chalk.red(`[Status: ${status}]`)
          );
        break;
    }

    console[level](logMessage);
  }

  info(message: string, status?: number) {
    this.log("info", message, status);
  }

  warn(message: string, status?: number) {
    this.log("warn", message, status);
  }

  error(message: string, status?: number) {
    this.log("error", message, status);
  }
}

const logger = new Logger();

export { logger };
