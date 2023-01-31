import '../styles/register.css';
import {useContext, useState} from "react";
import {BouncingButton} from "./common/bouncingButton";
import {registerUser} from "../api/user_api";
import {Navigate} from "react-router-dom";
import {NotificationContext} from "../contexts/notifContext";

export const Register = () => {
    const [email, setEmail] = useState('');
    const [password, setPassword] = useState('');
    const [passwordConfirmation, setPasswordConfirmation] = useState('');
    const [registrationCompleted, setRegistrationCompleted] = useState(false);
    const { addNotification } = useContext(NotificationContext);

    const handleSubmit = event => {
        event.preventDefault();

        if (!email || !password || !passwordConfirmation) {
            alert("Please fill in all the fields");
            return;
        }
        if (!/^\w+([.-]?\w+)*@\w+([.-]?\w+)*(\.\w{2,3})+$/.test(email)) {
            alert("Please enter a valid email address");
            return;
        }
        if (password.length < 8) {
            alert("Password should be at least 8 characters long");
            return;
        }
        if (password !== passwordConfirmation) {
            alert("Passwords do not match");
            return;
        }

        //handle result of registration
        registerUser({email, password}).then((r) => {
            r.ok ? setRegistrationCompleted(true) : r.json().then(r => addNotification(r.message))
        });
    };

    if (registrationCompleted) return <Navigate to={"/"}/>
    return (
        <div className="register">
            <form className="register" onSubmit={handleSubmit}>
                <input type="email" value={email} placeholder="Email" onChange={e => setEmail(e.target.value)} />
                <input type="password" value={password} placeholder="Password" onChange={e => setPassword(e.target.value)} />
                <input type="password" value={passwordConfirmation} placeholder="Confirm Password" onChange={e => setPasswordConfirmation(e.target.value)} />
                <BouncingButton className="register" type="submit">Register</BouncingButton>
            </form>
        </div>
    );
}