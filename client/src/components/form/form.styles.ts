import styled from 'styled-components';

export const FormContainer = styled.form`
    width: 600px;
    height: 60%;

    background-color: ${({ theme }) => theme.colors.background_focus};

    display: flex;
    align-items: center;
    justify-content: center;
    flex-direction: column;

    border-radius: 5px;
`;

export const FormHeader = styled.div`
    width: 100%;
    height: 50%;

    display: flex;
    align-items: flex-end;
    justify-content: center;
`

export const FormInputs = styled.div`
    width: 100%;
    min-height: 40px;

    display: flex;
    align-items: center;
    justify-content: center;
    flex-direction: column;
`


export const Input = styled.input`
    margin: 5px;
    padding: 0 10px;

    width: 500px;
    height: 50px;

    border-radius: 5px;
    background-color: ${({ theme }) => theme.colors.background};
    border: none;
`;
