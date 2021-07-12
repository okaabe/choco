import React from 'react'

import { User } from '../models/user'
import { usePersistedState } from '../helper/usePersistedState'
import { useEffect } from 'react'
import { rewoke } from '../services/auth'

export type AuthType = [User, {
    setToken?: React.Dispatch<React.SetStateAction<string>>,
    setUser?: React.Dispatch<React.SetStateAction<User>>
}]

export const AuthContext = React.createContext<AuthType>([{}, {}])

export const AuthHook: React.FC = ({
    children
}) => {
    const [user, setUser] = React.useState<User>({})
    const [token, setToken] = usePersistedState("choco@token", "")

    useEffect(() => {
        rewoke(token)
            .then((authentication) => {
                setUser({
                    jwt: token,
                    ...authentication,
                })
            }).catch((err) => {
                return console.log(err)
            })
    }, [token])

    return (
        <AuthContext.Provider value={[user, {
            setUser,
            setToken,
        }]}>
            { children }
        </AuthContext.Provider>
    )
}