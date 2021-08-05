import { MouseEventHandler, SetStateAction, Dispatch } from 'react';

export type FormInput = {
    id: string;
    className?: string;
    name?: string;
    placeholder?: string;
    isPassword?: boolean;
    stateDispatcher?: Dispatch<SetStateAction<string>>;
}

export type FormButton = {
    message?: string;
    onPress?: (event: MouseEventHandler<HTMLButtonElement>) => any;
}

export type FormProperties = {
    inputs: FormInput[];
    submitButton: FormButton;
    links?: string[];
}