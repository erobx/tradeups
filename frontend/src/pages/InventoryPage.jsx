import Inventory from "../components/Inventory/Inventory"

function InventoryPage() {
  return (
    <div className="flex flex-col items-center justify-center">
      <div className="flex">
        <h1 className="text-3xl font-bold">Inventory</h1>
      </div>
      <div className="mt-5">
        <Inventory />
      </div>
    </div>
  )
}

export default InventoryPage
