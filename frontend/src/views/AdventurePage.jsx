import React, { useState, useEffect, useMemo } from 'react'
import { useSearchParams } from 'react-router-dom'
import { capitalizeWords } from '../utils/formatting.js'
import './AdventurePage.css'

function AdventurePage() {
  const [searchParams] = useSearchParams()
  const [adventure, setAdventure] = useState(null)
  const [errorMsg, setErrorMsg] = useState('')
  const [isLoading, setIsLoading] = useState(false)
  const [isExpanded, setIsExpanded] = useState(false)

  const region = searchParams.get('region') || ''

  const buttonText = useMemo(() => {
    if (isLoading) return 'Finding Adventure...'
    return 'Pick Another Adventure'
  }, [isLoading])

  const displayName = useMemo(() => 
    capitalizeWords(adventure?.name), [adventure?.name]
  )

  const displayType = useMemo(() => 
    capitalizeWords(adventure?.type), [adventure?.type]
  )

  const getRandomAdventure = async () => {
    if (isLoading) return
    
    setIsLoading(true)
    setIsExpanded(false)
    
    try {
      const res = await fetch(`http://localhost:8080/random?region=${region}`)

      if (!res.ok) {
        const errorJson = await res.json()
        throw new Error(errorJson.details || errorJson.error || 'Something went wrong')
      }

      const data = await res.json()
      setAdventure(data)
      setErrorMsg('')
    } catch (err) {
      setAdventure(null)
      setErrorMsg(err.message)
    } finally {
      setIsLoading(false)
    }
  }

  useEffect(() => {
    getRandomAdventure()
  }, []) // eslint-disable-line react-hooks/exhaustive-deps

  return (
    <div id="app">
      <h1>Solo Adventure Picker</h1>

      {adventure && (
        <div className="card" key={adventure.name}>
          <div className="xp-badge">+{adventure.xpValue || 100}XP</div>
          <div className="card-content">
            <h2>{displayName}</h2>
            <h3>{displayType}</h3>
            
            <div className={`card-details ${isExpanded ? 'expanded' : ''}`}>
              <p className="description">
                {adventure.description || 'Embark on this mysterious adventure!'}
              </p>
              {adventure.tags?.length && (
                <div className="tags">
                  {adventure.tags.map(tag => (
                    <span key={tag} className="tag">{tag}</span>
                  ))}
                </div>
              )}
            </div>

            <button 
              className="toggle-details" 
              onClick={() => setIsExpanded(!isExpanded)}
            >
              {isExpanded ? 'Show Less' : 'Show More'}
            </button>
          </div>
        </div>
      )}

      {errorMsg && !adventure && (
        <div className="error card">
          {errorMsg}
        </div>
      )}

      <div className="reroll-btn-wrapper">
        <button 
          className="reroll-btn"
          onClick={getRandomAdventure} 
          disabled={isLoading}
        >
          {buttonText}
        </button>
      </div>
    </div>
  )
}

export default AdventurePage 