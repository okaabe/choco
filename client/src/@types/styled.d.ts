import 'styled-components';

declare module 'styled-components' {
    export interface DefaultTheme {
        name: string;
        colors: {
            background: string;
            background_focus: string;

            foreground: string;
            foreground_focus: string;

            widget_focus_color: string;
        }
    }
}
