import { BrowserRouter as Router, Routes, Route } from 'react-router-dom'
import ManagementScreen from '@/components/screens/ManagementScreen'

function App() {
  return (
    <Router>
      <Routes>
        <Route path="/" element={<ManagementScreen />} />
      </Routes>
    </Router>
  )
}

export default App
