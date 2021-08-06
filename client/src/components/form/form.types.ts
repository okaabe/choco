import React, { Dispatch, SetStateAction, DetailedHTMLProps } from 'react';

export type FormInput =  {
    placeholder?: string;
    type?: string;
    onChangeState?: Dispatch<SetStateAction<string>>;
}

export type FormProperties = {
    inputs?: FormInput[];
    links?: string[];
}