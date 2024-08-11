interface RequestOptions {
  method: string;
  body?: string;
  headers?: { [key: string]: string };
}
class FetchError extends Error {
  status: number;

  constructor(message: string, status: number) {
    super(message);
    this.name = "FetchError";
    this.status = status;
  }
}

export { FetchError };
export type { RequestOptions };
