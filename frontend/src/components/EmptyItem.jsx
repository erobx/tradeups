import plus from "../assets/plus-circle.svg"

function EmptyItem() {

  const addItem = () => {
    console.log("adding item...")
  }

  return (
    <div className="flex flex-col gap-2">
      <h1 className="text-lg">You currently have no skins in your inventory :(</h1>
      {/*
      <button
        className="btn btn-soft btn-success"
        onClick={addItem}
      >
        Add a skin to get started!
      </button>
      */}
    </div>
  )
}

export default EmptyItem
