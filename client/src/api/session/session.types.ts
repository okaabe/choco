import { User } from '../models/user';

export type AuthenticationCode = 
    'SUCESS'
    | 'FAILED'
    | 'MISSING_FIELDS'
    | 'UNAUTHORIZED';

export type AuthenticationResponse<T> = {
    code?: AuthenticationCode,
    message?: string,
    data?: T,
}