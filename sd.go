package iri

// http://www.w3.org/TR/sparql11-service-description/

// SPARQL 1.1 Service Description vocabulary.
const (
	// SPARQL 1.1 Service Description vocabulary.
	SD_NS = "http://www.w3.org/ns/sparql-service-description#"
	// Properties
	SD_endpoint                          = SD_NS + "endpoint"
	SD_feature                           = SD_NS + "feature"
	SD_defaultEntailmentRegime           = SD_NS + "defaultEntailmentRegime"
	SD_entailmentRegime                  = SD_NS + "entailmentRegime"
	SD_defaultSupportedEntailmentProfile = SD_NS + "defaultSupportedEntailmentProfile"
	SD_supportedEntailmentProfile        = SD_NS + "supportedEntailmentProfile"
	SD_extensionFunction                 = SD_NS + "extensionFunction"
	SD_extensionAggregate                = SD_NS + "extensionAggregate"
	SD_languageExtension                 = SD_NS + "languageExtension"
	SD_supportedLanguage                 = SD_NS + "supportedLanguage"
	SD_propertyFeature                   = SD_NS + "propertyFeature"
	SD_defaultDataset                    = SD_NS + "defaultDataset"
	SD_availableGraphs                   = SD_NS + "availableGraphs"
	SD_resultFormat                      = SD_NS + "resultFormat"
	SD_inputFormat                       = SD_NS + "inputFormat"
	SD_defaultGraph                      = SD_NS + "defaultGraph"
	SD_namedGraph                        = SD_NS + "namedGraph"
	SD_name                              = SD_NS + "name"
	SD_graph                             = SD_NS + "graph"
	// Classes
	SD_Service           = SD_NS + "Service"
	SD_Feature           = SD_NS + "Feature"
	SD_Language          = SD_NS + "Language"
	SD_Function          = SD_NS + "Function"
	SD_Aggregate         = SD_NS + "Aggregate"
	SD_EntailmentRegime  = SD_NS + "EntailmentRegime"
	SD_EntailmentProfile = SD_NS + "EntailmentProfile"
	SD_GraphCollection   = SD_NS + "GraphCollection"
	SD_Dataset           = SD_NS + "Dataset"
	SD_Graph             = SD_NS + "Graph"
	SD_NamedGraph        = SD_NS + "NamedGraph"
	// Instances
	SD_SPARQL10Query       = SD_NS + "SPARQL10Query"
	SD_SPARQL11Query       = SD_NS + "SPARQL11Query"
	SD_SPARQL11Update      = SD_NS + "SPARQL11Update"
	SD_DereferencesURIs    = SD_NS + "DereferencesURIs"
	SD_UnionDefaultGraph   = SD_NS + "UnionDefaultGraph"
	SD_RequiresDataset     = SD_NS + "RequiresDataset"
	SD_EmptyGraphs         = SD_NS + "EmptyGraphs"
	SD_BasicFederatedQuery = SD_NS + "BasicFederatedQuery"
)
