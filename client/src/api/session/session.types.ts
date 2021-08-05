export enum AuthenticationCode {
    BAD_REQUEST = 401,
    UNAUTHORIZED = 400,
    ACCEPTED = 202,
}

export type AuthenticationResponse<T> = {
    code?: AuthenticationCode,
    message?: string,
    data?: T,
}