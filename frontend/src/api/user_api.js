// user_api contains functions to communicate with backend api regarding user

// registerUser sends information about user that is trying to register to the backend
import {activateApi, registerApi} from "../const";

export const registerUser = async (user) => {
    const res = await fetch(registerApi, {
        method: 'POST',
        credentials: "include",
        body: JSON.stringify(user)})
    return res
}

export const activateUser = async (activationCode) => {
    const res = await fetch(activateApi + activationCode, {
        method: 'POST',
        credentials: "include",
    })
    const data = await res.json()

    return data.message
}
//
// export const loginUser = async (user: { email: string, password: string }) => {
//     const res = await fetch("http://localhost:8080/login", {
//         method: 'POST',
//         credentials: "include",
//         body: JSON.stringify(user)})
//     switch (res.status){
//         case 200:
//             return true
//         default:
//             showNotification({ message: 'Incorrect password', color: 'red' });
//             return false
//     }
// }
//
// export const logoutUser = async () => {
//     await fetch("http://localhost:8080/user/logout", {
//         credentials: "include",
//     })
//     showNotification({ message: 'User logged out', color: 'green' });
// }
//
// export const isLoggedIn = async () => {
//     const res = await fetch("http://localhost:8080/user/logged_in",{
//         credentials: "include"
//     })
//     return res.status === 200
// }
//
