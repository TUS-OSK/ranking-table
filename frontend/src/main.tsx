import { StrictMode } from 'react'
import { createRoot } from 'react-dom/client'
import Leaderboard from './leaderboard'
import './index.css'

createRoot(document.getElementById('root')!).render(
  <StrictMode>
    <Leaderboard />
  </StrictMode>,
)
