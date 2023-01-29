import '../styles/register.css';
import {useState} from "react";

export const Register = () => {
    const [email, setEmail] = useState('');
    const [password, setPassword] = useState('');
    const [passwordConfirmation, setPasswordConfirmation] = useState('');

    //TODO input validation

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
        // If all validations passed, send data to backend
        console.log(email, password, passwordConfirmation);
    };


    return (
        <div className="register">
            <form className="register" onSubmit={handleSubmit}>
                <input type="email" value={email} placeholder="Email" onChange={e => setEmail(e.target.value)} />
                <input type="password" value={password} placeholder="Password" onChange={e => setPassword(e.target.value)} />
                <input type="password" value={passwordConfirmation} placeholder="Confirm Password" onChange={e => setPasswordConfirmation(e.target.value)} />
                <button className="register" type="submit">Register</button>
            </form>
        </div>
    );
}