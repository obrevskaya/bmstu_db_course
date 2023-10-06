# GetWinesResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Wines** | [**[]Wine**](Wine.md) |  | 

## Methods

### NewGetWinesResponse

`func NewGetWinesResponse(wines []Wine, ) *GetWinesResponse`

NewGetWinesResponse instantiates a new GetWinesResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetWinesResponseWithDefaults

`func NewGetWinesResponseWithDefaults() *GetWinesResponse`

NewGetWinesResponseWithDefaults instantiates a new GetWinesResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetWines

`func (o *GetWinesResponse) GetWines() []Wine`

GetWines returns the Wines field if non-nil, zero value otherwise.

### GetWinesOk

`func (o *GetWinesResponse) GetWinesOk() (*[]Wine, bool)`

GetWinesOk returns a tuple with the Wines field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWines

`func (o *GetWinesResponse) SetWines(v []Wine)`

SetWines sets Wines field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


