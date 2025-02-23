import React, { useState } from 'react'


const Counter = () => {
    const [count, setCount] = useState(0)
    const handleClick = () => {
        //This will increment the count by 2
        setCount((prev) => prev + 1)
        setCount((prev) => prev + 1)
    }

    const handleReset = () => {
        //This will reset the count to 0
        setCount(0)
    }
    return (
        <div>
            <p>{count}</p>
            <button onClick={handleClick}>Click Me</button>
            <button onClick={handleReset}>Reset</button>
        </div>
    )
}

export default Counter