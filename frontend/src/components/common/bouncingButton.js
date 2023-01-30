import {useState} from "react";
import '../../styles/common/common.css';

// BouncingButton represents html button element that animates a bouncing movement when clicked
export const BouncingButton = ({ children, onClick }) => {
    const [clicked, setClicked] = useState(false);

    const handleClick = () => {
        setClicked(true);
        setTimeout(() => setClicked(false), 500);
        onClick && onClick();
    };

    return (
        <button
            className={ (clicked ? "common bouncing" : "common")}
            onClick={handleClick}
        >
            {children}
        </button>
    );
}
