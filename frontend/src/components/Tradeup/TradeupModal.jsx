import { useMemo, useState } from "react"
import ModalItem from "./ModalItem"
import useInventory from "../../stores/inventoryStore"

function TradeupModal({ tradeupId, rarity }) {
  const { inventory, setInventory, addItem, removeItem } = useInventory()
  const [currentPage, setCurrentPage] = useState(1)
  const [totalPages, setTotalPages] = useState(1)
  const itemsPerPage = 15

  const startIndex = (currentPage - 1) * itemsPerPage
  const endIndex = startIndex + itemsPerPage

  const paginate = (pageNumber) => setCurrentPage(pageNumber)

  const filtered = useMemo(() => {
    const temp = inventory.filter(i => i.rarity === rarity)
    const pages = Math.ceil(temp.length / itemsPerPage)
    setTotalPages(pages)

    return inventory.filter(i => i.rarity === rarity)
  }, [inventory])

  const currentItems = filtered.slice(startIndex, endIndex)

  return (
    <div className="h-48">
      <button className="btn btn-primary" onClick={()=>document.getElementById('modal_1').showModal()}>Add Skin</button>
      <dialog id="modal_1" className="modal">
        <div className="modal-box max-w-7xl max-h-3xl">
          <h3 className="font-bold text-lg mb-1">Showing all available skins...</h3>
          <div className="grid grid-cols-5 grid-rows-3 gap-2">
            {currentItems.map(s => (
              <ModalItem
                key={s.id}
                invId={s.id}
                tradeupId={tradeupId}
                name={s.name}
                wear={s.wear}
                price={s.price}
                isStatTrak={s.isStatTrak}
                imgSrc={s.imageSrc}
                removeItem={removeItem}
              />
            ))}
          </div>
          <div className="join mt-1">
            <button className="join-item btn" onClick={() => paginate(currentPage - 1)} disabled={currentPage === 1}>«</button>
            <button className="join-item btn">Page {currentPage}</button>
            <button className="join-item btn" onClick={() => paginate(currentPage + 1)} disabled={currentPage === totalPages}>»</button>
          </div>
        </div>
        <form method="dialog" className="modal-backdrop">
          <button>close</button>
        </form>
      </dialog>
    </div>
  )
}

export default TradeupModal
