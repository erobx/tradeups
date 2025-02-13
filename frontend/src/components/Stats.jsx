import { useEffect, useState } from "react"
import StatItem from "./StatItem"

function Stats() {
  //const [winnings, setWinnings] = useState([])
  const [entered, setEntered] = useState(50)
  const [won, setWon] = useState(10)
  const [earnings, setEarnings] = useState(32.41)

  const winnings = [
      {id: 0, name: "AK-47 | Slate", wear: "Factory New", rarity: "Classified", isStatTrak: true, imgSrc: "/ak-slate.png", price: 37.21},
      {id: 0, name: "Five-SeveN | Candy Apple", wear: "Minimal Wear", rarity: "Restricted", isStatTrak: false, imgSrc: "/fs-candy-apple.png", price: 15.32},
      {id: 0, name: "Five-SeveN | Candy Apple", wear: "Minimal Wear", rarity: "Restricted", isStatTrak: false, imgSrc: "/fs-candy-apple.png", price: 15.32},
      {id: 0, name: "AK-47 | Slate", wear: "Factory New", rarity: "Classified", isStatTrak: true, imgSrc: "/ak-slate.png", price: 37.21},
    ]
  
  return (
    <div className="flex flex-col items-center">
      <h1 className="font-bold text-info text-3xl mb-0.5">Statistics</h1>
      <div className="stats bg-base-300 shadow-md">
        <div className="stat">
          <div className="stat-title">Recent winnings</div>
          <div className="flex">
            {winnings.map((skin, index) => (
              <StatItem
                name={skin.name}
                wear={skin.wear}
                rarity={skin.rarity}
                isStatTrak={skin.isStatTrak}
                imgSrc={skin.imgSrc}
                price={skin.price}
              />
            ))}
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
