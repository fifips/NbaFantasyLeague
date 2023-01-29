import './styles/App.css';
import {Header} from './components/header/header'
import {BrowserRouter, Route, Routes} from "react-router-dom";
import {Register} from "./components/register";

const App = () => {
  return (
      <div className="App">
        <Header />
        <BrowserRouter>
            <Routes>
                <Route path="/register" element={<Register/>}/>
            </Routes>
        </BrowserRouter>
      </div>
  );
}

export default App;