import axios from 'axios'

import { User } from '../models/user'
import { BASE_API_URL } from './api'


export async function signin(email: string, password: string): Promise<User | null> {
    const response = await axios.post(`${BASE_API_URL}/api/auth/signin`, {
        email, password
    })

    if (response.data.err) {
        throw new Error(response.data.err)
    }

    return response.data
}

export async function rewoke(token: string): Promise<Omit<User, "jwt"> | null> {
    const response = await axios.get(`${BASE_API_URL}/api/auth/rewoke`, {
        headers: {
            "Authorization": token,
        }
    })

    if (response.data.err) {
        throw new Error(response.data.err)
    }

    return response.data
}

export async function signup(username: string, email: string, password: string): Promise<User | null> {
    const response = await axios.post(`${BASE_API_URL}/api/auth/signup`, {
        username, email, password,
    })

    if (response.data.err) {
        throw new Error(response.data.err)
    }

    return response.data
}
