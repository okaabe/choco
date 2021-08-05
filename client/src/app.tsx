import React from 'react';

import { View } from './components/view';
import { Toast } from './components/popup/toast';

import { ApplicationRoutes } from './router';

import { SessionProvider } from './context/session';
import { ApplicationThemeProvider } from './context/themes';

import { Dark } from './styles/themes/dark';

export const App: React.FC = () => {
    return (
        <View>
            <ApplicationThemeProvider defaultTheme={ Dark }>
                <SessionProvider>
                    <ApplicationRoutes />
                    <Toast />
                </SessionProvider>
            </ApplicationThemeProvider>
        </View>
    );
}
