import { BrowserRouter as Router, Routes, Route } from 'react-router-dom';
import SignInSide from './componets/SignInSide';
import LoginVarProvider from './componets/logindetails';
import SignUp from './componets/SignUp';
import AllPosts from './componets/AllPosts';

function App() {
  return (
    <LoginVarProvider>
      <Router>
        <Routes>
          <Route path='/' element={<AllPosts />}/>
          <Route path="/signin" element={<SignInSide />} />
          <Route path="/signup" element={<SignUp />} />
        </Routes>
      </Router>
     </LoginVarProvider>
  );
}

export default App;

// useEffect(() => {  }, [process.env.REACT_APP_LoggedIn]);