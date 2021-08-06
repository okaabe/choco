import React from 'react';

import { Toast } from './components/popup/toast';
import { ApplicationRoutes } from './router';
import { ApplicationThemeProvider } from './context/themes';
import { ApplicationStyles } from './styles/globalStyled';

import { Light } from './styles/themes/light';

export const App: React.FC = () => {
    return (
        <ApplicationThemeProvider defaultTheme={ Light }>
            <ApplicationRoutes />
            <Toast />
            <ApplicationStyles />
        </ApplicationThemeProvider>
    );
}
