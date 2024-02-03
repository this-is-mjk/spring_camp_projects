import { BrowserRouter as Router, Routes, Route } from 'react-router-dom';
import SignInSide from './componets/SignInSide';
import SignUp from './componets/SignUp';
import AllPosts from './componets/AllPosts';

function App() {
  return (
    <Router>
      <Routes>
        <Route path='/' element={<AllPosts />}/>
        <Route path="/signin" element={<SignInSide />} />
        <Route path="/signup" element={<SignUp />} />
      </Routes>
    </Router>
  );
}

export default App;
