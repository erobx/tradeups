import { useEffect, useState } from "react"
import StatItem from "./StatItem"
import { getStats } from "../../api/user"

function Stats({ user }) {
  const [winnings, setWinnings] = useState([])
  const [entered, setEntered] = useState(0)
  const [won, setWon] = useState(0)
  const [earnings, setEarnings] = useState(0)
  
  const fetchStats = async () => {
    const jwt = localStorage.getItem("jwt")
    try {
      const data = await getStats(jwt, user.id)
      if (data) {
        setEntered(data.tradeupsEntered)
        setWon(data.tradeupsWon)
        if (data.recentWinnings) {
          setWinnings(data.recentWinnings)
        }
      }
    } catch (error) {
      console.error("Error fetching stats: ", error)
    }
  }

  useEffect(() => {
    fetchStats()
  }, [])
  
  return (
    <div className="flex flex-col items-center">
      <div className="stats bg-base-300 shadow-md">
        <div className="stat">
          <div className="stat-title">Recent winnings</div>
          <div className="flex">
            {winnings.length === 0 ? (
              <div>No winnings</div>
            ) : (
              winnings.map((skin, index) => (
                <StatItem
                  key={index}
                  name={skin.name}
                  wear={skin.wear}
                  rarity={skin.rarity}
                  isStatTrak={skin.isStatTrak}
                  imgSrc={skin.imageSrc}
                  price={skin.price}
                />
              ))
            )}
          </div>
        </div>
      </div>

      <div className="divider"></div>

      <div className="stats bg-base-300 shadow-md">
        <div className="stat place-items-center">
          <div className="stat-title">Trade Ups Entered</div>
          <div className="stat-value text-primary">{entered}</div>
        </div>

        <div className="stat place-items-center">
          <div className="stat-title">Trade Ups Won</div>
          <div className="stat-value text-accent">{won}</div>
        </div>

        <div className="stat place-items-center">
          <div className="stat-title">Total Earnings</div>
          <div className="stat-value text-primary">${earnings}</div>
        </div>
      </div>
    </div>
  )
}

export default Stats
