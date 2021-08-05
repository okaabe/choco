import axios from 'axios';

import { BASE_API_URL } from '../api';

import { User } from '../models/user';
import { AuthenticationResponse } from './session.types';

export async function authenticate(email: string, password: string): Promise<AuthenticationResponse<User>> {
    return {};
}

export async function create(username: string, email: string, password: string): Promise<AuthenticationResponse<User>> {
    return {};
}

export async function rewoke(token: string): Promise<AuthenticationResponse<User>> {
    return {};
}
