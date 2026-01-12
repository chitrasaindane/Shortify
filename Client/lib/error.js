const defaultErrMsg = "Something Went Wrong!";

// # Send 'Error'
const sendError = (errMsg) => {
    if (!errMsg) {
        errMsg = defaultErrMsg;
    }

    const err = new Error(errMsg);
    throw err;
};

export {
    sendError
};