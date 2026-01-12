const defaultSuccessMsg = "Operation Successful!";

// # Get 'Success Message'
const getSuccessMsg = (successMsg) => {
    if (!successMsg) {
        successMsg = defaultSuccessMsg;
    }

    return successMsg;
};

export {
    getSuccessMsg
};