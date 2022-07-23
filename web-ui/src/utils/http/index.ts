export class HttpError extends Error {
  info: object;
  status: number;
  constructor(message: string, info: object, status: number) {
    super(message);
    this.info = info;
    this.status = status;
  }
}
