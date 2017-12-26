/*
Package DOM provides type definitions for use with the Chrome DOM protocol

https://chromedevtools.github.io/devtools-protocol/tot/DOM/
*/
package DOM

import (
	"github.com/mkenney/go-chrome/protocol/page"
	"github.com/mkenney/go-chrome/protocol/runtime"
)

/*
CollectClassNamesFromSubtreeParams represents DOM.collectClassNamesFromSubtree parameters.

https://chromedevtools.github.io/devtools-protocol/tot/DOM/#method-collectClassNamesFromSubtree
*/
type CollectClassNamesFromSubtreeParams struct {
	// ID of the node to collect class names.
	NodeID NodeID `json:"nodeId"`
}

/*
CollectClassNamesFromSubtreeResult represents the result of calls to
DOM.collectClassNamesFromSubtree.

https://chromedevtools.github.io/devtools-protocol/tot/DOM/#method-collectClassNamesFromSubtree
*/
type CollectClassNamesFromSubtreeResult struct {
	// Class name list.
	ClassNames []string `json:"classNames"`
}

/*
CopyToParams represents DOM.copyTo parameters.

https://chromedevtools.github.io/devtools-protocol/tot/DOM/#method-copyTo
*/
type CopyToParams struct {
	// ID of the node to copy.
	NodeID NodeID `json:"nodeId"`

	// ID of the element to drop the copy into.
	TargetNodeID NodeID `json:"targetNodeId"`

	// Optional. Drop the copy before this node (if absent, the copy becomes the last child of
	// targetNodeId).
	InsertBeforeNodeID NodeID `json:"insertBeforeNodeId,omitempty"`
}

/*
CopyToResult represents the result of calls to DOM.copyTo.

https://chromedevtools.github.io/devtools-protocol/tot/DOM/#method-copyTo
*/
type CopyToResult struct {
	// ID of the new cloned node.
	NodeID NodeID `json:"nodeId"`
}

/*
DescribeNodeParams represents DOM.describeNode parameters.

https://chromedevtools.github.io/devtools-protocol/tot/DOM/#method-describeNode
*/
type DescribeNodeParams struct {
	// Optional. ID of the node.
	NodeID NodeID `json:"nodeId,omitempty"`

	// Optional. ID of the backend node.
	BackendNodeID BackendNodeID `json:"backendNodeId,omitempty"`

	// Optional. JavaScript object ID of the node wrapper.
	ObjectID Runtime.RemoteObjectID `json:"objectId,omitempty"`

	// Optional. The maximum depth at which children should be retrieved, defaults to 1. Use -1 for
	// the entire subtree or provide an integer larger than 0.
	Depth int `json:"depth,omitempty"`

	// Optional. Whether or not iframes and shadow roots should be traversed when returning the
	// subtree (default is false).
	Pierce bool `json:"pierce,omitempty"`
}

/*
DescribeNodeResult represents the result of calls to DOM.describeNode.

https://chromedevtools.github.io/devtools-protocol/tot/DOM/#method-describeNode
*/
type DescribeNodeResult struct {
	// ID of the new cloned node.
	NodeID NodeID `json:"nodeId"`
}

/*
DiscardSearchResultsParams represents DOM.discardSearchResults parameters.

https://chromedevtools.github.io/devtools-protocol/tot/DOM/#method-discardSearchResults
*/
type DiscardSearchResultsParams struct {
	// Node description.
	Node Node `json:"node"`
}

/*
FocusParams represents DOM.focus parameters.

https://chromedevtools.github.io/devtools-protocol/tot/DOM/#method-focus
*/
type FocusParams struct {
	// Optional. ID of the node.
	NodeID NodeID `json:"nodeId,omitempty"`

	// Optional. ID of the backend node.
	BackendNodeID BackendNodeID `json:"backendNodeId,omitempty"`

	// Optional. JavaScript object ID of the node wrapper.
	ObjectID Runtime.RemoteObjectID `json:"objectId,omitempty"`
}

/*
GetAttributesParams represents DOM.getAttributes parameters.

https://chromedevtools.github.io/devtools-protocol/tot/DOM/#method-getAttributes
*/
type GetAttributesParams struct {
	// ID  of the node to retrieve attibutes for.
	NodeID NodeID `json:"nodeId"`
}

/*
GetAttributesResult represents the result of calls to DOM.getAttributes.

https://chromedevtools.github.io/devtools-protocol/tot/DOM/#method-getAttributes
*/
type GetAttributesResult struct {
	// An interleaved array of node attribute names and values.
	Attributes []string `json:"attributes"`
}

/*
GetBoxModelParams represents DOM.getBoxModel parameters.

https://chromedevtools.github.io/devtools-protocol/tot/DOM/#method-getBoxModel
*/
type GetBoxModelParams struct {
	// Optional. ID of the node.
	NodeID NodeID `json:"nodeId,omitempty"`

	// Optional. ID of the backend node.
	BackendNodeID BackendNodeID `json:"backendNodeId,omitempty"`

	// Optional. JavaScript object ID of the node wrapper.
	ObjectID Runtime.RemoteObjectID `json:"objectId,omitempty"`
}

/*
GetBoxModelResult represents the result of calls to DOM.getBoxModel.

https://chromedevtools.github.io/devtools-protocol/tot/DOM/#method-getBoxModel
*/
type GetBoxModelResult struct {
	// Box model for the node.
	Model BoxModel `json:"model"`
}

/*
GetDocumentParams represents DOM.getDocument parameters.

https://chromedevtools.github.io/devtools-protocol/tot/DOM/#method-getDocument
*/
type GetDocumentParams struct {
	// Optional. The maximum depth at which children should be retrieved, defaults to 1. Use -1 for
	// the entire subtree or provide an integer larger than 0.
	Depth int `json:"depth,omitempty"`

	// Optional. Whether or not iframes and shadow roots should be traversed when returning the
	// subtree (default is false).
	Pierce bool `json:"pierce,omitempty"`
}

/*
GetDocumentResult represents the result of calls to DOM.getDocument.

https://chromedevtools.github.io/devtools-protocol/tot/DOM/#method-getDocument
*/
type GetDocumentResult struct {
	// Resulting node.
	Root Node `json:"root"`
}

/*
GetFlattenedDocumentParams represents DOM.getFlattenedDocument parameters.

https://chromedevtools.github.io/devtools-protocol/tot/DOM/#method-getFlattenedDocument
*/
type GetFlattenedDocumentParams struct {
	// Optional. The maximum depth at which children should be retrieved, defaults to 1. Use -1 for
	// the entire subtree or provide an integer larger than 0.
	Depth int `json:"depth,omitempty"`

	// Optional. Whether or not iframes and shadow roots should be traversed when returning the
	// subtree (default is false).
	Pierce bool `json:"pierce,omitempty"`
}

/*
GetFlattenedDocumentResult represents the result of calls to DOM.getFlattenedDocument.

https://chromedevtools.github.io/devtools-protocol/tot/DOM/#method-getFlattenedDocument
*/
type GetFlattenedDocumentResult struct {
	// Resulting nodes.
	Nodes []Node `json:"nodes"`
}

/*
GetNodeForLocationParams represents DOM.getNodeForLocation parameters.

https://chromedevtools.github.io/devtools-protocol/tot/DOM/#method-getNodeForLocation
*/
type GetNodeForLocationParams struct {
	// X coordinate.
	X int `json:"x"`

	// Y coordinate.
	Y int `json:"y"`

	// Optional. False to skip to the nearest non-UA shadow root ancestor (default: false).
	IncludeUserAgentShadowDOM bool `json:"includeUserAgentShadowDOM,omitempty"`
}

/*
GetNodeForLocationResult represents the result of calls to DOM.getNodeForLocation.

https://chromedevtools.github.io/devtools-protocol/tot/DOM/#method-getNodeForLocation
*/
type GetNodeForLocationResult struct {
	// ID of the node at given coordinates.
	NodeID NodeID `json:"nodeId"`
}

/*
GetOuterHTMLParams represents DOM.getOuterHTML parameters.

https://chromedevtools.github.io/devtools-protocol/tot/DOM/#method-getOuterHTML
*/
type GetOuterHTMLParams struct {
	// Optional. ID of the node.
	NodeID NodeID `json:"nodeId,omitempty"`

	// Optional. ID of the backend node.
	BackendNodeID BackendNodeID `json:"backendNodeId,omitempty"`

	// Optional. JavaScript object ID of the node wrapper.
	ObjectID Runtime.RemoteObjectID `json:"objectId,omitempty"`
}

/*
GetOuterHTMLResult represents the result of calls to DOM.getOuterHTML.

https://chromedevtools.github.io/devtools-protocol/tot/DOM/#method-getOuterHTML
*/
type GetOuterHTMLResult struct {
	// Outer HTML markup.
	OuterHTML string `json:"outerHTML"`
}

/*
GetRelayoutBoundaryParams represents DOM.getRelayoutBoundary parameters.

https://chromedevtools.github.io/devtools-protocol/tot/DOM/#method-getRelayoutBoundary
*/
type GetRelayoutBoundaryParams struct {
	// ID of the node.
	NodeID NodeID `json:"nodeId"`
}

/*
GetRelayoutBoundaryResult represents the result of calls to DOM.getRelayoutBoundary.

https://chromedevtools.github.io/devtools-protocol/tot/DOM/#method-getRelayoutBoundary
*/
type GetRelayoutBoundaryResult struct {
	// Relayout boundary node ID for the given node.
	NodeID NodeID `json:"nodeId"`
}

/*
GetSearchResultsParams represents DOM.getSearchResults parameters.

https://chromedevtools.github.io/devtools-protocol/tot/DOM/#method-getSearchResults
*/
type GetSearchResultsParams struct {
	// Unique search session identifier.
	SearchID string `json:"searchId"`

	// Start index of the search result to be returned.
	FromIndex int `json:"fromIndex"`

	// End index of the search result to be returned.
	ToIndex int `json:"toIndex"`
}

/*
GetSearchResultsResult represents the result of calls to DOM.getSearchResults.

https://chromedevtools.github.io/devtools-protocol/tot/DOM/#method-getSearchResults
*/
type GetSearchResultsResult struct {
	// IDs of the search result nodes.
	NodeIDs []NodeID `json:"nodeIds"`
}

/*
MoveToParams represents DOM.moveTo parameters.

https://chromedevtools.github.io/devtools-protocol/tot/DOM/#method-moveTo
*/
type MoveToParams struct {
	// ID of the node to move.
	NodeID NodeID `json:"nodeId"`

	// ID of the element to drop the moved node into.
	TargetNodeID NodeID `json:"targetNodeId"`

	// Optional. Drop node before this one (if absent, the moved node becomes the last child of
	// targetNodeId).
	InsertBeforeNodeID NodeID `json:"insertBeforeNodeId,omitempty"`
}

/*
MoveToResult represents the result of calls to DOM.moveTo.

https://chromedevtools.github.io/devtools-protocol/tot/DOM/#method-moveTo
*/
type MoveToResult struct {
	// New ID of the moved node.
	NodeID NodeID `json:"nodeId"`
}

/*
PerformSearchParams represents DOM.performSearch parameters.

https://chromedevtools.github.io/devtools-protocol/tot/DOM/#method-performSearch
*/
type PerformSearchParams struct {
	// Plain text or query selector or XPath search query.
	Query string `json:"query"`

	// Optional. True to search in user agent shadow DOM.
	IncludeUserAgentShadowDOM bool `json:"includeUserAgentShadowDOM,omitempty"`
}

/*
PerformSearchResult represents the result of calls to DOM.performSearch.

https://chromedevtools.github.io/devtools-protocol/tot/DOM/#method-performSearch
*/
type PerformSearchResult struct {
	// Unique search session identifier.
	SearchID string `json:"searchId"`

	// Number of search results.
	ResultCount int `json:"resultCount"`
}

/*
PushNodeByPathToFrontendParams represents DOM.pushNodeByPathToFrontend parameters.

https://chromedevtools.github.io/devtools-protocol/tot/DOM/#method-pushNodeByPathToFrontend
*/
type PushNodeByPathToFrontendParams struct {
	// Path to node in the proprietary format.
	Path string `json:"path"`
}

/*
PushNodeByPathToFrontendResult represents the result of calls to DOM.pushNodeByPathToFrontend.

https://chromedevtools.github.io/devtools-protocol/tot/DOM/#method-pushNodeByPathToFrontend
*/
type PushNodeByPathToFrontendResult struct {
	// ID of the node for given path.
	NodeID NodeID `json:"nodeId"`
}

/*
PushNodesByBackendIDsToFrontendParams represents DOM.pushNodesByBackendIdsToFrontend parameters.

https://chromedevtools.github.io/devtools-protocol/tot/DOM/#method-pushNodesByBackendIdsToFrontend
*/
type PushNodesByBackendIDsToFrontendParams struct {
	// The array of backend node IDs.
	BackendNodeIDs []BackendNodeID `json:"backendNodeIds"`
}

/*
PushNodesByBackendIDsToFrontendResult represents the result of calls to
DOM.pushNodesByBackendIdsToFrontend.

https://chromedevtools.github.io/devtools-protocol/tot/DOM/#method-pushNodesByBackendIdsToFrontend
*/
type PushNodesByBackendIDsToFrontendResult struct {
	// The array of IDs of pushed nodes that correspond to the backend IDs specified in
	// backendNodeIDs.
	NodeIDs []NodeID `json:"nodeIds"`
}

/*
QuerySelectorParams represents DOM.querySelector parameters.

https://chromedevtools.github.io/devtools-protocol/tot/DOM/#method-querySelector
*/
type QuerySelectorParams struct {
	// ID of the node to query.
	NodeID NodeID `json:"nodeId"`

	// Selector string.
	Selector string `json:"selector"`
}

/*
QuerySelectorResult represents the result of calls to DOM.querySelector.

https://chromedevtools.github.io/devtools-protocol/tot/DOM/#method-querySelector
*/
type QuerySelectorResult struct {
	// Query selector result.
	NodeID NodeID `json:"nodeId"`
}

/*
QuerySelectorAllParams represents DOM.querySelectorAll parameters.

https://chromedevtools.github.io/devtools-protocol/tot/DOM/#method-querySelectorAll
*/
type QuerySelectorAllParams struct {
	// ID of the node to query upon.
	NodeID NodeID `json:"nodeId"`

	// Selector string.
	Selector string `json:"selector"`
}

/*
QuerySelectorAllResult represents the result of calls to DOM.querySelectorAll.

https://chromedevtools.github.io/devtools-protocol/tot/DOM/#method-querySelectorAll
*/
type QuerySelectorAllResult struct {
	// Query selector result.
	NodeIDs []NodeID `json:"nodeIds"`
}

/*
RemoveAttributeParams represents DOM.removeAttribute parameters.

https://chromedevtools.github.io/devtools-protocol/tot/DOM/#method-removeAttribute
*/
type RemoveAttributeParams struct {
	// ID of the element to remove attribute from.
	NodeID NodeID `json:"nodeId"`

	// Name of the attribute to remove.
	Name string `json:"name"`
}

/*
RemoveNodeParams represents DOM.removeNode parameters.

https://chromedevtools.github.io/devtools-protocol/tot/DOM/#method-removeNode
*/
type RemoveNodeParams struct {
	// ID of the node to remove.
	NodeID NodeID `json:"nodeId"`
}

/*
RequestChildNodesParams represents DOM.requestChildNodes parameters.

https://chromedevtools.github.io/devtools-protocol/tot/DOM/#method-requestChildNodes
*/
type RequestChildNodesParams struct {
	// ID of the node to get children for.
	NodeID NodeID `json:"nodeId"`

	// Optional. The maximum depth at which children should be retrieved, defaults to 1. Use -1 for
	// the entire subtree or provide an integer larger than 0.
	Depth int `json:"depth,omitempty"`

	// Optional. Whether or not iframes and shadow roots should be traversed when returning the
	// sub-tree (default is false).
	Pierce bool `json:"pierce,omitempty"`
}

/*
RequestNodeParams represents DOM.requestNode parameters.

https://chromedevtools.github.io/devtools-protocol/tot/DOM/#method-requestNode
*/
type RequestNodeParams struct {
	// JavaScript object ID to convert into node.
	ObjectID Runtime.RemoteObjectID `json:"objectId"`
}

/*
RequestNodeResult represents the result of calls to DOM.requestNode.

https://chromedevtools.github.io/devtools-protocol/tot/DOM/#method-requestNode
*/
type RequestNodeResult struct {
	// Node ID for given object.
	NodeID NodeID `json:"nodeId"`
}

/*
ResolveNodeParams represents DOM.resolveNode parameters.

https://chromedevtools.github.io/devtools-protocol/tot/DOM/#method-resolveNode
*/
type ResolveNodeParams struct {
	// Optional. ID of the node to resolve.
	NodeID NodeID `json:"nodeId,omitempty"`

	// Optional. Backend identifier of the node to resolve.
	BackendNodeID BackendNodeID `json:"backendNodeId,omitempty"`

	// Optional. Symbolic group name that can be used to release multiple objects.
	ObjectGroup string `json:"objectGroup,omitempty"`
}

/*
ResolveNodeResult represents the result of calls to DOM.resolveNode.

https://chromedevtools.github.io/devtools-protocol/tot/DOM/#method-resolveNode
*/
type ResolveNodeResult struct {
	// JavaScript object wrapper for given node.
	Object Runtime.RemoteObject `json:"object"`
}

/*
SetAttributeValueParams represents DOM.setAttributeValue parameters.

https://chromedevtools.github.io/devtools-protocol/tot/DOM/#method-setAttributeValue
*/
type SetAttributeValueParams struct {
	// ID of the element to set attribute for.
	NodeID NodeID `json:"nodeId"`

	// Attribute name.
	Name string `json:"name"`

	// Attribute value.
	Value string `json:"value"`
}

/*
SetAttributesAsTextParams represents DOM.setAttributesAsText parameters.

https://chromedevtools.github.io/devtools-protocol/tot/DOM/#method-setAttributesAsText
*/
type SetAttributesAsTextParams struct {
	// ID of the element to set attributes for.
	NodeID NodeID `json:"nodeId"`

	// Text with a number of attributes. Will parse this text using HTML parser.
	Text string `json:"text"`

	// Optional. Attribute name to replace with new attributes derived from text in case text parsed
	// successfully.
	Name string `json:"name,omitempty"`
}

/*
SetFileInputFilesParams represents DOM.setFileInputFiles parameters.

https://chromedevtools.github.io/devtools-protocol/tot/DOM/#method-setFileInputFiles
*/
type SetFileInputFilesParams struct {
	// Array of file paths to set.
	Files []string `json:"files"`

	// Optional. ID of the node.
	NodeID NodeID `json:"nodeId,omitempty"`

	// Optional. ID of the backend node.
	BackendNodeID BackendNodeID `json:"backendNodeId,omitempty"`

	// Optional. JavaScript object ID of the node wrapper.
	ObjectID Runtime.RemoteObjectID `json:"objectId,omitempty"`
}

/*
SetInspectedNodeParams represents DOM.setInspectedNode parameters.

https://chromedevtools.github.io/devtools-protocol/tot/DOM/#method-setInspectedNode
*/
type SetInspectedNodeParams struct {
	// DOM node ID to be accessible by means of $x command line API.
	NodeID NodeID `json:"nodeId"`
}

/*
SetNodeNameParams represents DOM.setNodeName parameters.

https://chromedevtools.github.io/devtools-protocol/tot/DOM/#method-setNodeName
*/
type SetNodeNameParams struct {
	// ID of the node to set name for.
	NodeID NodeID `json:"nodeId"`

	// New node name.
	Name string `json:"name"`
}

/*
SetNodeNameResult represents the result of calls to DOM.setNodeName.

https://chromedevtools.github.io/devtools-protocol/tot/DOM/#method-setNodeName
*/
type SetNodeNameResult struct {
	// New node's ID.
	NodeID NodeID `json:"nodeId"`
}

/*
SetNodeValueParams represents DOM.setNodeValue parameters.

https://chromedevtools.github.io/devtools-protocol/tot/DOM/#method-setNodeValue
*/
type SetNodeValueParams struct {
	// ID of the node to set value for.
	NodeID NodeID `json:"nodeId"`

	// New node value.
	Value string `json:"value"`
}

/*
SetOuterHTMLParams represents DOM.setOuterHTML parameters.

https://chromedevtools.github.io/devtools-protocol/tot/DOM/#method-setOuterHTML
*/
type SetOuterHTMLParams struct {
	// ID of the node to set markup for.
	NodeID NodeID `json:"nodeId"`

	// Outer HTML markup to set.
	OuterHTML string `json:"outerHTML"`
}

/*
AttributeModifiedEvent represents DOM.attributeModified event data.

https://chromedevtools.github.io/devtools-protocol/tot/DOM/#event-attributeModified
*/
type AttributeModifiedEvent struct {
	// ID of the node that has changed.
	NodeID NodeID `json:"nodeId"`

	// Attribute name.
	Name string `json:"name"`

	// Attribute value.
	Value string `json:"value"`
}

/*
AttributeRemovedEvent represents DOM.attributeRemoved event data.

https://chromedevtools.github.io/devtools-protocol/tot/DOM/#event-attributeRemoved
*/
type AttributeRemovedEvent struct {
	// ID of the node that has changed.
	NodeID NodeID `json:"nodeId"`

	// Attribute name.
	Name string `json:"name"`
}

/*
CharacterDataModifiedEvent represents DOM.characterDataModified event data.

https://chromedevtools.github.io/devtools-protocol/tot/DOM/#event-characterDataModified
*/
type CharacterDataModifiedEvent struct {
	// ID of the node that has changed.
	NodeID NodeID `json:"nodeId"`

	// New text value.
	CharacterData string `json:"characterData"`
}

/*
ChildNodeCountUpdatedEvent represents DOM.childNodeCountUpdated event data.

https://chromedevtools.github.io/devtools-protocol/tot/DOM/#event-childNodeCountUpdated
*/
type ChildNodeCountUpdatedEvent struct {
	// ID of the node that has changed.
	ParentNodeID NodeID `json:"parentNodeId"`

	// If of the previous siblint.
	PreviousNodeID NodeID `json:"previousNodeId"`

	// Inserted node data.
	Node Node `json:"node"`
}

/*
ChildNodeInsertedEvent represents DOM.childNodeInserted event data.

https://chromedevtools.github.io/devtools-protocol/tot/DOM/#event-childNodeInserted
*/
type ChildNodeInsertedEvent struct {
	// ID of the node that has changed.
	ParentNodeID NodeID `json:"parentNodeId"`

	// If of the previous siblint.
	PreviousNodeID NodeID `json:"previousNodeId"`

	// Inserted node data.
	Node Node `json:"node"`
}

/*
ChildNodeRemovedEvent represents DOM.childNodeRemoved event data.

https://chromedevtools.github.io/devtools-protocol/tot/DOM/#event-childNodeRemoved
*/
type ChildNodeRemovedEvent struct {
	// Parent ID.
	ParentNodeID NodeID `json:"parentNodeId"`

	// ID of the node that has been removed.
	NodeID NodeID `json:"nodeId"`
}

/*
DistributedNodesUpdatedEvent represents DOM.distributedNodesUpdated event data.

https://chromedevtools.github.io/devtools-protocol/tot/DOM/#event-distributedNodesUpdated
*/
type DistributedNodesUpdatedEvent struct {
	// Insertion point where distributed nodes were updated.
	InsertionPointID NodeID `json:"insertionPointId"`

	// Distributed nodes for given insertion point.
	DistributedNodes []*BackendNode `json:"distributedNodes"`
}

/*
DocumentUpdatedEvent represents DOM.documentUpdated event data.

https://chromedevtools.github.io/devtools-protocol/tot/DOM/#event-documentUpdated
*/
type DocumentUpdatedEvent struct{}

/*
InlineStyleInvalidatedEvent represents DOM.inlineStyleInvalidated event data.

https://chromedevtools.github.io/devtools-protocol/tot/DOM/#event-inlineStyleInvalidated
*/
type InlineStyleInvalidatedEvent struct {
	// IDs of the nodes for which the inline styles have been invalidated.
	NodeIDs []NodeID `json:"nodeIds"`
}

/*
PseudoElementAddedEvent represents DOM.pseudoElementAdded event data.

https://chromedevtools.github.io/devtools-protocol/tot/DOM/#event-pseudoElementAdded
*/
type PseudoElementAddedEvent struct {
	// Pseudo element's parent element ID.
	ParentID NodeID `json:"parentId"`

	// The added pseudo element.
	PseudoElement Node `json:"pseudoElement"`
}

/*
PseudoElementRemovedEvent represents DOM.pseudoElementRemoved event data.

https://chromedevtools.github.io/devtools-protocol/tot/DOM/#event-pseudoElementRemoved
*/
type PseudoElementRemovedEvent struct {
	// Pseudo element's parent element ID.
	ParentID NodeID `json:"parentId"`

	// The removed pseudo element ID.
	PseudoElementID NodeID `json:"pseudoElementId"`
}

/*
SetChildNodesEvent represents DOM.setChildNodes event data.

https://chromedevtools.github.io/devtools-protocol/tot/DOM/#event-setChildNodes
*/
type SetChildNodesEvent struct {
	// Parent node ID to populate with children.
	ParentID NodeID `json:"parentId"`

	// Child nodes array.
	Nodes []*Node `json:"nodes"`
}

/*
ShadowRootPoppedEvent represents DOM.shadowRootPopped event data.

https://chromedevtools.github.io/devtools-protocol/tot/DOM/#event-shadowRootPopped
*/
type ShadowRootPoppedEvent struct {
	// Host element ID.
	HostID NodeID `json:"hostId"`

	// Shadow root ID.
	RootID NodeID `json:"rootId"`
}

/*
ShadowRootPushedEvent represents DOM.shadowRootPushed event data.

https://chromedevtools.github.io/devtools-protocol/tot/DOM/#event-shadowRootPushed
*/
type ShadowRootPushedEvent struct {
	// Host element ID.
	HostID NodeID `json:"hostId"`

	// Shadow root.
	Root Node `json:"root"`
}

/*
NodeID is a unique DOM node identifier.

https://chromedevtools.github.io/devtools-protocol/tot/DOM/#type-NodeId
*/
type NodeID int

/*
BackendNodeID is a unique DOM node identifier used to reference a node that may not have been pushed
to the front-end.

https://chromedevtools.github.io/devtools-protocol/tot/DOM/#type-BackendNodeId
*/
type BackendNodeID int

/*
BackendNode is a backend node with a friendly name.

https://chromedevtools.github.io/devtools-protocol/tot/DOM/#type-BackendNode
*/
type BackendNode struct {
	// Node's nodeType.
	NodeType int `json:"nodeType"`

	// Node's nodeName.
	NodeName string `json:"nodeName"`

	// Node's ID
	BackendNodeID BackendNodeID `json:"backendNodeId"`
}

/*
PseudoType is a pseudo element type.

https://chromedevtools.github.io/devtools-protocol/tot/DOM/#type-PseudoType
*/
type PseudoType string

/*
ShadowRootType is a shadow root type.

https://chromedevtools.github.io/devtools-protocol/tot/DOM/#type-ShadowRootType
*/
type ShadowRootType string

/*
Node is a base node mirror type. DOM interaction is implemented in terms of mirror objects that
represent the actual DOM nodes.

https://chromedevtools.github.io/devtools-protocol/tot/DOM/#type-Node
*/
type Node struct {
	// Node identifier that is passed into the rest of the DOM messages as the nodeId. Backend will
	// only push node with given id once. It is aware of all requested nodes and will only fire DOM
	// events for nodes known to the client.
	NodeID NodeID `json:"nodeId"`

	// Optional. The ID of the parent node if any.
	ParentID NodeID `json:"parentId,omitempty"`

	// The BackendNodeID for this node.
	BackendNodeID BackendNodeID `json:"backendNodeId"`

	// Node's nodeType.
	NodeType int `json:"nodeType"`

	// Node's nodeName.
	NodeName string `json:"nodeName"`

	// Node's localName.
	LocalName string `json:"localName"`

	// Node's nodeValue.
	NodeValue string `json:"nodeValue"`

	// Optional. Child count for Container nodes.
	ChildNodeCount int `json:"childNodeCount,omitempty"`

	// Optional. Child nodes of this node when requested with children.
	Children []*Node `json:"children,omitempty"`

	// Optional. Attributes of the Element node in the form of flat array
	// [name1, value1, name2, value2].
	Attributes []string `json:"attributes,omitempty"`

	// Optional. Document URL that Document or FrameOwner node points to.
	DocumentURL string `json:"documentURL,omitempty"`

	// Optional. Base URL that Document or FrameOwner node uses for URL
	// completion.
	BaseURL string `json:"baseURL,omitempty"`

	// Optional. DocumentType's publicId.
	PublicID string `json:"publicId,omitempty"`

	// Optional. DocumentType's systemId.
	SystemID string `json:"systemId,omitempty"`

	// Optional. DocumentType's internalSubset.
	InternalSubset string `json:"internalSubset,omitempty"`

	// Optional. Document's XML version in case of XML documents.
	XMLVersion string `json:"xmlVersion,omitempty"`

	// Optional. Attr's name.
	Name string `json:"name,omitempty"`

	// Optional. Attr's value.
	Value string `json:"value,omitempty"`

	// Optional. Pseudo element type for this node.
	PseudoType PseudoType `json:"pseudoType,omitempty"`

	// Optional. Shadow root type.
	ShadowRootType ShadowRootType `json:"shadowRootType,omitempty"`

	// Optional. Frame ID for frame owner elements.
	FrameID Page.FrameID `json:"frameId,omitempty"`

	// Optional. Content document for frame owner elements.
	ContentDocument *Node `json:"contentDocument,omitempty"`

	// Optional. Shadow root list for given element host.
	ShadowRoots []*Node `json:"shadowRoots,omitempty"`

	// Optional. Content document fragment for template elements.
	TemplateContent *Node `json:"templateContent,omitempty"`

	// Optional. Pseudo elements associated with this node.
	PseudoElements []*Node `json:"pseudoElements,omitempty"`

	// Optional. Import document for the HTMLImport links.
	ImportedDocument *Node `json:"importedDocument,omitempty"`

	// Optional. Distributed nodes for given insertion point.
	DistributedNodes []*BackendNode `json:"distributedNodes,omitempty"`

	// Optional. Whether the node is SVG.
	IsSVG bool `json:"isSVG,omitempty"`
}

/*
RGBA is a structure holding an RGBA color.

https://chromedevtools.github.io/devtools-protocol/tot/DOM/#type-RGBA
*/
type RGBA struct {
	// The red component, in the [0-255] range.
	R int `json:"r"`

	// The green component, in the [0-255] range.
	G int `json:"g"`

	// The blue component, in the [0-255] range.
	B int `json:"b"`

	// Optional. The alpha component, in the [0-1] range (default: 1).
	A int `json:"a,omitempty"`
}

/*
Quad is an array of quad vertices, x immediately followed by y for each point, points clock-wise.

https://chromedevtools.github.io/devtools-protocol/tot/DOM/#type-Quad
*/
type Quad [2]int

/*
BoxModel represents the box model.

https://chromedevtools.github.io/devtools-protocol/tot/DOM/#type-BoxModel
*/
type BoxModel struct {
	// Content box.
	Content Quad `json:"content"`

	// Padding box.
	Padding Quad `json:"padding"`

	// Border box.
	Border Quad `json:"border"`

	// Margin box.
	Margin Quad `json:"margin"`

	// Node width.
	Width int `json:"width"`

	// Node height.
	Height int `json:"height"`

	// Optional. Shape outside coordinates.
	//
	// This expects an instance of ShapeOutsideInfo, but that doesn't omitempty correctly so it must
	// be added manually.
	ShapeOutside interface{} `json:"shapeOutside,omitempty"`
}

/*
ShapeOutsideInfo represents the CSS Shape Outside details.

https://chromedevtools.github.io/devtools-protocol/tot/DOM/#type-ShapeOutsideInfo
*/
type ShapeOutsideInfo struct {
	// Shape bounds.
	Bounds Quad `json:"bounds"`

	// Shape coordinate details.
	Shape []interface{} `json:"shape"`

	// Margin shape bounds.
	MarginShape []interface{} `json:"marginShape"`
}

/*
Rect defines a rectangle

https://chromedevtools.github.io/devtools-protocol/tot/DOM/#type-Rect
*/
type Rect struct {
	// X coordinate.
	X float64 `json:"x"`

	// Y coordinate.
	Y float64 `json:"y"`

	// Rectangle width.
	Width float64 `json:"width"`

	// Rectangle height.
	Height float64 `json:"height"`
}