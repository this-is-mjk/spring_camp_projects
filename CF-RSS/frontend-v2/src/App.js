import Login from './components/Login';
import Signup from './components/Signup';
import AllPostBox from './components/AllPostBox';
import { BrowserRouter, Routes, Route } from "react-router-dom";
import AllCommentsPage from './components/AllCommentsPage';
import Navbar from './components/NavBar';
import ProfilePage from './components/ProfilePage';
function App() {
  localStorage.clear();
  return (
    <BrowserRouter>
      <Routes>
        <Route path="/" element={<><Navbar /> <AllPostBox /></>} />
        <Route path="/login" element={<Login />} />
        <Route path="/signup" element={<Signup />} />
        <Route path="/AllCommentsPage" element ={<><Navbar /> <AllCommentsPage /> </>} />
        <Route path="/myprofile" element ={<><Navbar /> <ProfilePage /> </>} />
      </Routes>
    </BrowserRouter>
  );
}
export default App;
