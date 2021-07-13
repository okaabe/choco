import styled from 'styled-components';

export const SignUpContainer = styled.div`
    width: 100vw;
    height: 100vh;
    overflow: hidden;
`;

export const SignUpContent = styled.div`
    width: 100%;
    height: 100%;

    display: flex;
    align-items: center;
`;

export const SignUpForm = styled.form`
    width: 35%;
    height: 100%;
    
    background-color: #fafafa;

    display: flex;
    align-items: center;
    justify-content: center;
    flex-direction: column;
`;

export const SignFormMessage = styled.div`
    width: 500px;

    h1 {
        font-size: 30px;
    }
    p {
        margin: 10px 0 10px 0;
        text-align: left;
        letter-spacing: 1px;
    }
`;

export const SignUpFormInputs = styled.div`
    width: 100%;

    display: flex;
    flex-direction: column;
    justify-content: center;
    align-items: center;

    input {
        width: 500px;
        height: 45px;
        background-color: #ffffff;
        border-radius: 5px;
        border-style: solid;
        border-color: #EAEAEA;

        padding: 0 15px 0 15px;
        margin: 5px;
    }
`;

export const SignUpThumbnail = styled.div`
    width: 65%;
    height: 100%;
    overflow: hidden;

    display: flex;
    align-items: flex-end;
    flex-direction: row-reverse;

    img {
        width: 100%;

    }
`;

export const SignUpButton = styled.button`
    margin: 10px;

    background-color: #F48023;
    opacity: .5;

    width: 500px;
    height: 47px;

    border: none;

    transition: opacity .5s;
    cursor: pointer;

    &:hover {
        opacity: 1;
    }
`;

export const SignUpErrorMessage = styled.div`
    color: #FF0000;
    margin: 5px 0 5px 0;
`