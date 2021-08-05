import React from 'react';

import { View } from './components/view';
import { Toast } from './components/popup/toast';

import { ApplicationRoutes } from './router';

import { SessionProvider } from './context/session';
import { ApplicationThemeProvider } from './context/themes';

import { Dark } from './styles/themes/dark';

export const App: React.FC = () => {
    return (
        <ApplicationThemeProvider defaultTheme={ Dark }>
            <ApplicationRoutes />
            <Toast />
        </ApplicationThemeProvider>
    );
}
