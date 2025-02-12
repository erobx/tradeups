import { badgeMap } from "../constants/rarity"

function RarityBadge({ rarity }) {
 const badgeColor = badgeMap[rarity]

  return (
    <span className={`badge ${badgeColor}`}>{rarity}</span>
  )
}

export default RarityBadge
