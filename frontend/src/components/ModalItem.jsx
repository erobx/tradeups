import StatTrakBadge from "./StatTrakBadge"

function ModalItem({ name, wear, isStatTrak, imgSrc, price }) {
  return (
    <div className="card card-md w-48 h-48 bg-base-300">
      <h1 className="text-sm font-bold text-primary ml-1.5 mt-0.5">${price}</h1>
      <figure>
        <img
          alt={imgSrc}
          src={imgSrc}
          width={100}
          height={50}
        />
      </figure>
      <div className="card-body items-center">
        <h1 className="card-title text-sm">{name}</h1>
        <h1 className="card-title text-xs">({wear})</h1>
        <div className="flex gap-2">
          {isStatTrak && <StatTrakBadge />}
        </div>
      </div>
    </div>
  )
}

export default ModalItem
