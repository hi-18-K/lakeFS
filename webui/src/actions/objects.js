import * as api from "./api";
import {AsyncActionType} from "./request";
import { Link,  Element, Events, scrollSpy, scroller, animateScroll as scroll } from "react-scroll";
import * as Scroll from 'react-scroll';

export const PAGINATION_AMOUNT = 300;
export const OBJECTS_LIST_TREE = new AsyncActionType('OBJECTS_GET_TREE');
export const OBJECTS_LIST_TREE_PAGINATE = new AsyncActionType('OBJECTS_GET_TREE_PAGINATE');
export const OBJECTS_UPLOAD = new AsyncActionType('OBJECTS_UPLOAD');
export const OBJECTS_DELETE = new AsyncActionType('OBJECTS_DELETE');


export const listTree = (repoId, branchId, tree, amount = PAGINATION_AMOUNT, readUncommitted = true) => {
    return OBJECTS_LIST_TREE.execute(async () => {
        return await api.objects.list(repoId, branchId, tree, "", amount, readUncommitted);
    });
};

export const listTreePaginate = (repoId, branchId, tree, after, amount = PAGINATION_AMOUNT, readUncommitted = true) => {
    return OBJECTS_LIST_TREE_PAGINATE.execute(async () => {
        return await api.objects.list(repoId, branchId, tree, after, amount, readUncommitted);
    });
};


export const upload = (repoId, branchId, path, fileObject) => {
    return OBJECTS_UPLOAD.execute(async () => {
        return await api.objects.upload(repoId, branchId, path, fileObject);
    });
};

export const uploadDone = () => {
    return OBJECTS_UPLOAD.resetAction();
};

let yoffset = 0;

export const deleteObject = (repoId, branchId, path) => {
    constructor(props) {
      super(props);
      this.state = {yoffset: 0};
    }
    storeScroll = () => {
      console.log(window.pageYOffset);
      yoffset = window.pageYOffset;
      this.setState({yoffset: window.pageYOffset});
    };
    return OBJECTS_DELETE.execute(async () => {
        this.storeScroll;
        return await api.objects.delete(repoId, branchId, path);
    });
};

export const deleteObjectDone = () => {
    scrollToPageOffset = offset_val => e =>  {
      scroll.scrollTo(offset_val);
    };
    return (
      OBJECTS_DELETE.resetAction();
      this.scrollToPageOffset(yoffset);
      yoffset = 0;
    )
};
