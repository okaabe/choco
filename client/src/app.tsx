import React from 'react'

import { Styles } from './styles/styled'
import { AuthHook } from './hooks/auth'
import { AppRoutes } from './routes';

export const App: React.FC = () => {
    return (
        <>
        <AuthHook>
            <AppRoutes />
        </AuthHook>
        <Styles/>
        </>
    )
}