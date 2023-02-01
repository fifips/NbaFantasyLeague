import "../../styles/navigationMenu/navigationMenu.css"

export const NavigationMenu = () => {
        return (
            <nav className="nav-menu">
                <ul>
                    <li><a href="/">Home</a></li>
                    <li><a href="/leagues">Leagues</a></li>
                    <li><a href="/">Services</a></li>
                    <li><a href="/">Contact</a></li>
                </ul>
            </nav>
        );
}