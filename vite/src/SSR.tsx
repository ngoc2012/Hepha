import { Routes, Route } from 'react-router-dom';

import Home from './pages/home.tsx';
import About from './pages/about.tsx';


function SSR() {
  return (
    <>
      <Routes>
         <Route path="/" element={<Home />} />
         <Route path="/about" element={<About />} />
      </Routes>
    </>
  )
}

export default SSR