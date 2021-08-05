import React from 'react';

import { View } from './components/view';
import { Toast } from './components/popup/toast';

import { ApplicationRoutes } from './router';

import { SessionProvider } from './context/session';

export const App: React.FC = () => {
    return (
        <View>
            <SessionProvider>
                <ApplicationRoutes />
                <Toast />
            </SessionProvider>
        </View>
    )
}