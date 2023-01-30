import "../../styles/common/closeButton.css"

// CloseButton is a button displayed as an X, representing closing/removing action
export const CloseButton = ({ className, onClick }) => {
    return (
        <button className={className + " close"} onClick={onClick}/>
    );
}