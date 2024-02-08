import React,  { createContext, useContext, useState }  from "react";

const Logindetail = createContext();

export const useGlobalVar = () => useContext(Logindetail);
const LoginVarProvider = (props) => {
    const [emailid, setEmail] = useState("");
    const [password, setPassword] = useState("");
    const [isLoggedIn, setLogin] = useState(false);


    const updateEmail = (newEmail) => {
    setEmail(newEmail);
    };

    const updatePassword = (newPassword) => {
    setPassword(newPassword);
    };

    const updateIsLoggedIn = (newIsLoggedIn) => {
    setLogin(newIsLoggedIn);
    };
    return (
        <Logindetail.Provider value={{
            emailid,
            password,
            isLoggedIn,
            updateEmail,
            updatePassword,
            updateIsLoggedIn,

        }}>
            {props.children}
        </Logindetail.Provider>
    )
}
export default LoginVarProvider