// Code generated by smithy-go-codegen DO NOT EDIT.

package types

import (
	"time"
)

// Represents the history of a specific alarm.
type AlarmHistoryItem struct {
	// Data about the alarm, in JSON format.
	HistoryData *string
	// The time stamp for the alarm history item.
	Timestamp *time.Time
	// A summary of the alarm history, in text format.
	HistorySummary *string
	// The type of alarm history item.
	HistoryItemType HistoryItemType
	// The descriptive name for the alarm.
	AlarmName *string
	// The type of alarm, either metric alarm or composite alarm.
	AlarmType AlarmType
}

// An anomaly detection model associated with a particular CloudWatch metric and
// statistic. You can use the model to display a band of expected normal values
// when the metric is graphed.
type AnomalyDetector struct {
	// The current status of the anomaly detector's training. The possible values are
	// TRAINED | PENDING_TRAINING | TRAINED_INSUFFICIENT_DATA
	StateValue AnomalyDetectorStateValue
	// The metric dimensions associated with the anomaly detection model.
	Dimensions []*Dimension
	// The namespace of the metric associated with the anomaly detection model.
	Namespace *string
	// The name of the metric associated with the anomaly detection model.
	MetricName *string
	// The statistic associated with the anomaly detection model.
	Stat *string
	// The configuration specifies details about how the anomaly detection model is to
	// be trained, including time ranges to exclude from use for training the model,
	// and the time zone to use for the metric.
	Configuration *AnomalyDetectorConfiguration
}

// The configuration specifies details about how the anomaly detection model is to
// be trained, including time ranges to exclude from use for training the model and
// the time zone to use for the metric.
type AnomalyDetectorConfiguration struct {
	// The time zone to use for the metric. This is useful to enable the model to
	// automatically account for daylight savings time changes if the metric is
	// sensitive to such time changes. To specify a time zone, use the name of the time
	// zone as specified in the standard tz database. For more information, see tz
	// database (https://en.wikipedia.org/wiki/Tz_database).
	MetricTimezone *string
	// An array of time ranges to exclude from use when the anomaly detection model is
	// trained. Use this to make sure that events that could cause unusual values for
	// the metric, such as deployments, aren't used when CloudWatch creates the model.
	ExcludedTimeRanges []*Range
}

// The details about a composite alarm.
type CompositeAlarm struct {
	// The rule that this alarm uses to evaluate its alarm state.
	AlarmRule *string
	// An explanation for the alarm state, in JSON format.
	StateReasonData *string
	// The state value for the alarm.
	StateValue StateValue
	// The Amazon Resource Name (ARN) of the alarm.
	AlarmArn *string
	// The actions to execute when this alarm transitions to the INSUFFICIENT_DATA
	// state from any other state. Each action is specified as an Amazon Resource Name
	// (ARN).
	InsufficientDataActions []*string
	// Indicates whether actions should be executed during any changes to the alarm
	// state.
	ActionsEnabled *bool
	// An explanation for the alarm state, in text format.
	StateReason *string
	// The name of the alarm.
	AlarmName *string
	// The actions to execute when this alarm transitions to the OK state from any
	// other state. Each action is specified as an Amazon Resource Name (ARN).
	OKActions []*string
	// The time stamp of the last update to the alarm configuration.
	AlarmConfigurationUpdatedTimestamp *time.Time
	// The time stamp of the last update to the alarm state.
	StateUpdatedTimestamp *time.Time
	// The actions to execute when this alarm transitions to the ALARM state from any
	// other state. Each action is specified as an Amazon Resource Name (ARN).
	AlarmActions []*string
	// The description of the alarm.
	AlarmDescription *string
}

// Represents a specific dashboard.
type DashboardEntry struct {
	// The size of the dashboard, in bytes.
	Size *int64
	// The time stamp of when the dashboard was last modified, either by an API call or
	// through the console. This number is expressed as the number of milliseconds
	// since Jan 1, 1970 00:00:00 UTC.
	LastModified *time.Time
	// The name of the dashboard.
	DashboardName *string
	// The Amazon Resource Name (ARN) of the dashboard.
	DashboardArn *string
}

// An error or warning for the operation.
type DashboardValidationMessage struct {
	// A message describing the error or warning.
	Message *string
	// The data path related to the message.
	DataPath *string
}

// Encapsulates the statistical data that CloudWatch computes from metric data.
type Datapoint struct {
	// The minimum metric value for the data point.
	Minimum *float64
	// The standard unit for the data point.
	Unit StandardUnit
	// The sum of the metric values for the data point.
	Sum *float64
	// The time stamp used for the data point.
	Timestamp *time.Time
	// The percentile statistic for the data point.
	ExtendedStatistics map[string]*float64
	// The number of metric values that contributed to the aggregate value of this data
	// point.
	SampleCount *float64
	// The maximum metric value for the data point.
	Maximum *float64
	// The average of the metric values that correspond to the data point.
	Average *float64
}

// A dimension is a name/value pair that is part of the identity of a metric. You
// can assign up to 10 dimensions to a metric. Because dimensions are part of the
// unique identifier for a metric, whenever you add a unique name/value pair to one
// of your metrics, you are creating a new variation of that metric.
type Dimension struct {
	// The value of the dimension.
	Value *string
	// The name of the dimension. Dimension names cannot contain blank spaces or
	// non-ASCII characters.
	Name *string
}

// Represents filters for a dimension.
type DimensionFilter struct {
	// The value of the dimension to be matched.
	Value *string
	// The dimension name to be matched.
	Name *string
}

// This structure contains the definition for a Contributor Insights rule.
type InsightRule struct {
	// Indicates whether the rule is enabled or disabled.
	State *string
	// For rules that you create, this is always {"Name": "CloudWatchLogRule",
	// "Version": 1}. For built-in rules, this is {"Name": "ServiceLogRule", "Version":
	// 1}
	Schema *string
	// The name of the rule.
	Name *string
	// The definition of the rule, as a JSON object. The definition contains the
	// keywords used to define contributors, the value to aggregate on if this rule
	// returns a sum instead of a count, and the filters. For details on the valid
	// syntax, see Contributor Insights Rule Syntax
	// (https://docs.aws.amazon.com/AmazonCloudWatch/latest/monitoring/ContributorInsights-RuleSyntax.html).
	Definition *string
}

// One of the unique contributors found by a Contributor Insights rule. If the rule
// contains multiple keys, then a unique contributor is a unique combination of
// values from all the keys in the rule. If the rule contains a single key, then
// each unique contributor is each unique value for this key. For more information,
// see GetInsightRuleReport
// (https://docs.aws.amazon.com/AmazonCloudWatch/latest/APIReference/API_GetInsightRuleReport.html).
type InsightRuleContributor struct {
	// One of the log entry field keywords that is used to define contributors for this
	// rule.
	Keys []*string
	// An array of the data points where this contributor is present. Only the data
	// points when this contributor appeared are included in the array.
	Datapoints []*InsightRuleContributorDatapoint
	// An approximation of the aggregate value that comes from this contributor.
	ApproximateAggregateValue *float64
}

// One data point related to one contributor. For more information, see
// GetInsightRuleReport
// (https://docs.aws.amazon.com/AmazonCloudWatch/latest/APIReference/API_GetInsightRuleReport.html)
// and InsightRuleContributor
// (https://docs.aws.amazon.com/AmazonCloudWatch/latest/APIReference/API_InsightRuleContributor.html).
type InsightRuleContributorDatapoint struct {
	// The approximate value that this contributor added during this timestamp.
	ApproximateValue *float64
	// The timestamp of the data point.
	Timestamp *time.Time
}

// One data point from the metric time series returned in a Contributor Insights
// rule report. For more information, see GetInsightRuleReport
// (https://docs.aws.amazon.com/AmazonCloudWatch/latest/APIReference/API_GetInsightRuleReport.html).
type InsightRuleMetricDatapoint struct {
	// The minimum value from a single contributor during the time period represented
	// by that data point. This statistic is returned only if you included it in the
	// Metrics array in your request.
	Minimum *float64
	// The number of occurrences that matched the rule during this data point. This
	// statistic is returned only if you included it in the Metrics array in your
	// request.
	SampleCount *float64
	// The number of unique contributors who published data during this timestamp. This
	// statistic is returned only if you included it in the Metrics array in your
	// request.
	UniqueContributors *float64
	// The maximum value from a single occurence from a single contributor during the
	// time period represented by that data point. This statistic is returned only if
	// you included it in the Metrics array in your request.
	Maximum *float64
	// The average value from all contributors during the time period represented by
	// that data point. This statistic is returned only if you included it in the
	// Metrics array in your request.
	Average *float64
	// The timestamp of the data point.
	Timestamp *time.Time
	// The sum of the values from all contributors during the time period represented
	// by that data point. This statistic is returned only if you included it in the
	// Metrics array in your request.
	Sum *float64
	// The maximum value provided by one contributor during this timestamp. Each
	// timestamp is evaluated separately, so the identity of the max contributor could
	// be different for each timestamp. This statistic is returned only if you included
	// it in the Metrics array in your request.
	MaxContributorValue *float64
}

// A message returned by the GetMetricDataAPI, including a code and a description.
type MessageData struct {
	// The message text.
	Value *string
	// The error code or status code associated with the message.
	Code *string
}

// Represents a specific metric.
type Metric struct {
	// The namespace of the metric.
	Namespace *string
	// The name of the metric. This is a required field.
	MetricName *string
	// The dimensions for the metric.
	Dimensions []*Dimension
}

// The details about a metric alarm.
type MetricAlarm struct {
	// The description of the alarm.
	AlarmDescription *string
	// The state value for the alarm.
	StateValue StateValue
	// The actions to execute when this alarm transitions to the INSUFFICIENT_DATA
	// state from any other state. Each action is specified as an Amazon Resource Name
	// (ARN).
	InsufficientDataActions []*string
	// The name of the metric associated with the alarm, if this is an alarm based on a
	// single metric.
	MetricName *string
	// In an alarm based on an anomaly detection model, this is the ID of the
	// ANOMALY_DETECTION_BAND function used as the threshold for the alarm.
	ThresholdMetricId *string
	// The actions to execute when this alarm transitions to the ALARM state from any
	// other state. Each action is specified as an Amazon Resource Name (ARN).
	AlarmActions []*string
	// The name of the alarm.
	AlarmName *string
	// The period, in seconds, over which the statistic is applied.
	Period *int32
	// The number of periods over which data is compared to the specified threshold.
	EvaluationPeriods *int32
	// An explanation for the alarm state, in JSON format.
	StateReasonData *string
	// The time stamp of the last update to the alarm configuration.
	AlarmConfigurationUpdatedTimestamp *time.Time
	// An explanation for the alarm state, in text format.
	StateReason *string
	// The Amazon Resource Name (ARN) of the alarm.
	AlarmArn *string
	// The dimensions for the metric associated with the alarm.
	Dimensions []*Dimension
	// Indicates whether actions should be executed during any changes to the alarm
	// state.
	ActionsEnabled *bool
	// The namespace of the metric associated with the alarm.
	Namespace *string
	// The percentile statistic for the metric associated with the alarm. Specify a
	// value between p0.0 and p100.
	ExtendedStatistic *string
	// Used only for alarms based on percentiles. If ignore, the alarm state does not
	// change during periods with too few data points to be statistically significant.
	// If evaluate or this parameter is not used, the alarm is always evaluated and
	// possibly changes state no matter how many data points are available.
	EvaluateLowSampleCountPercentile *string
	// Sets how this alarm is to handle missing data points. If this parameter is
	// omitted, the default behavior of missing is used.
	TreatMissingData *string
	// The time stamp of the last update to the alarm state.
	StateUpdatedTimestamp *time.Time
	// The number of data points that must be breaching to trigger the alarm.
	DatapointsToAlarm *int32
	// The unit of the metric associated with the alarm.
	Unit StandardUnit
	// The arithmetic operation to use when comparing the specified statistic and
	// threshold. The specified statistic value is used as the first operand.
	ComparisonOperator ComparisonOperator
	// The value to compare with the specified statistic.
	Threshold *float64
	// The actions to execute when this alarm transitions to the OK state from any
	// other state. Each action is specified as an Amazon Resource Name (ARN).
	OKActions []*string
	// An array of MetricDataQuery structures, used in an alarm based on a metric math
	// expression. Each structure either retrieves a metric or performs a math
	// expression.  One item in the Metrics array is the math expression that the alarm
	// watches. This expression by designated by having <code>ReturnValue</code> set to
	// true.</p>
	Metrics []*MetricDataQuery
	// The statistic for the metric associated with the alarm, other than percentile.
	// For percentile statistics, use ExtendedStatistic.
	Statistic Statistic
}

// This structure is used in both GetMetricData and PutMetricAlarm. The supported
// use of this structure is different for those two operations. When used in
// GetMetricData, it indicates the metric data to return, and whether this call is
// just retrieving a batch set of data for one metric, or is performing a math
// expression on metric data. A single GetMetricData call can include up to 500
// MetricDataQuery structures. When used in PutMetricAlarm, it enables you to
// create an alarm based on a metric math expression. Each MetricDataQuery in the
// array specifies either a metric to retrieve, or a math expression to be
// performed on retrieved metrics. A single PutMetricAlarm call can include up to
// 20 MetricDataQuery structures in the array. The 20 structures can include as
// many as 10 structures that contain a MetricStat parameter to retrieve a metric,
// and as many as 10 structures that contain the Expression parameter to perform a
// math expression. Of those Expression structures, one must have True as the value
// for ReturnData. The result of this expression is the value the alarm watches.
// <p>Any expression used in a <code>PutMetricAlarm</code> operation must return a
// single time series. For more information, see <a
// href="https://docs.aws.amazon.com/AmazonCloudWatch/latest/monitoring/using-metric-math.html#metric-math-syntax">Metric
// Math Syntax and Functions</a> in the <i>Amazon CloudWatch User Guide</i>.</p>
// <p>Some of the parameters of this structure also have different uses whether you
// are using this structure in a <code>GetMetricData</code> operation or a
// <code>PutMetricAlarm</code> operation. These differences are explained in the
// following parameter list.</p>
type MetricDataQuery struct {
	// The math expression to be performed on the returned data, if this object is
	// performing a math expression. This expression can use the Id of the other
	// metrics to refer to those metrics, and can also use the Id of other expressions
	// to use the result of those expressions. For more information about metric math
	// expressions, see Metric Math Syntax and Functions
	// (https://docs.aws.amazon.com/AmazonCloudWatch/latest/monitoring/using-metric-math.html#metric-math-syntax)
	// in the Amazon CloudWatch User Guide. Within each MetricDataQuery object, you
	// must specify either Expression or MetricStat but not both.
	Expression *string
	// The granularity, in seconds, of the returned data points. For metrics with
	// regular resolution, a period can be as short as one minute (60 seconds) and must
	// be a multiple of 60. For high-resolution metrics that are collected at intervals
	// of less than one minute, the period can be 1, 5, 10, 30, 60, or any multiple of
	// 60. High-resolution metrics are those metrics stored by a PutMetricData
	// operation that includes a StorageResolution of 1 second.
	Period *int32
	// When used in GetMetricData, this option indicates whether to return the
	// timestamps and raw data values of this metric. If you are performing this call
	// just to do math expressions and do not also need the raw data returned, you can
	// specify False. If you omit this, the default of True is used. When used in
	// PutMetricAlarm, specify True for the one expression result to use as the alarm.
	// For all other metrics and expressions in the same PutMetricAlarm operation,
	// specify ReturnData as False.
	ReturnData *bool
	// A short name used to tie this object to the results in the response. This name
	// must be unique within a single call to GetMetricData. If you are performing math
	// expressions on this set of data, this name represents that data and can serve as
	// a variable in the mathematical expression. The valid characters are letters,
	// numbers, and underscore. The first character must be a lowercase letter.
	Id *string
	// The metric to be returned, along with statistics, period, and units. Use this
	// parameter only if this object is retrieving a metric and not performing a math
	// expression on returned data. Within one MetricDataQuery object, you must specify
	// either Expression or MetricStat but not both.
	MetricStat *MetricStat
	// A human-readable label for this metric or expression. This is especially useful
	// if this is an expression, so that you know what the value represents. If the
	// metric or expression is shown in a CloudWatch dashboard widget, the label is
	// shown. If Label is omitted, CloudWatch generates a default.
	Label *string
}

// A GetMetricData call returns an array of MetricDataResult structures. Each of
// these structures includes the data points for that metric, along with the
// timestamps of those data points and other identifying information.
type MetricDataResult struct {
	// The status of the returned data. Complete indicates that all data points in the
	// requested time range were returned. PartialData means that an incomplete set of
	// data points were returned. You can use the NextToken value that was returned and
	// repeat your request to get more data points. NextToken is not returned if you
	// are performing a math expression. InternalError indicates that an error
	// occurred. Retry your request using NextToken, if present.
	StatusCode StatusCode
	// A list of messages with additional information about the data returned.
	Messages []*MessageData
	// The timestamps for the data points, formatted in Unix timestamp format. The
	// number of timestamps always matches the number of values and the value for
	// Timestamps[x] is Values[x].
	Timestamps []*time.Time
	// The short name you specified to represent this metric.
	Id *string
	// The human-readable label associated with the data.
	Label *string
	// The data points for the metric corresponding to Timestamps. The number of values
	// always matches the number of timestamps and the timestamp for Values[x] is
	// Timestamps[x].
	Values []*float64
}

// Encapsulates the information sent to either create a metric or add new values to
// be aggregated into an existing metric.
type MetricDatum struct {
	// When you are using a Put operation, this defines what unit you want to use when
	// storing the metric. In a Get operation, this displays the unit that is used for
	// the metric.
	Unit StandardUnit
	// Array of numbers representing the values for the metric during the period. Each
	// unique value is listed just once in this array, and the corresponding number in
	// the Counts array specifies the number of times that value occurred during the
	// period. You can include up to 150 unique values in each PutMetricData action
	// that specifies a Values array. Although the Values array accepts numbers of type
	// Double, CloudWatch rejects values that are either too small or too large. Values
	// must be in the range of -2^360 to 2^360. In addition, special values (for
	// example, NaN, +Infinity, -Infinity) are not supported.
	Values []*float64
	// Valid values are 1 and 60. Setting this to 1 specifies this metric as a
	// high-resolution metric, so that CloudWatch stores the metric with sub-minute
	// resolution down to one second. Setting this to 60 specifies this metric as a
	// regular-resolution metric, which CloudWatch stores at 1-minute resolution.
	// Currently, high resolution is available only for custom metrics. For more
	// information about high-resolution metrics, see High-Resolution Metrics
	// (https://docs.aws.amazon.com/AmazonCloudWatch/latest/monitoring/publishingMetrics.html#high-resolution-metrics)
	// in the Amazon CloudWatch User Guide. This field is optional, if you do not
	// specify it the default of 60 is used.
	StorageResolution *int32
	// The dimensions associated with the metric.
	Dimensions []*Dimension
	// The time the metric data was received, expressed as the number of milliseconds
	// since Jan 1, 1970 00:00:00 UTC.
	Timestamp *time.Time
	// The name of the metric.
	MetricName *string
	// The statistical values for the metric.
	StatisticValues *StatisticSet
	// Array of numbers that is used along with the Values array. Each number in the
	// Count array is the number of times the corresponding value in the Values array
	// occurred during the period. If you omit the Counts array, the default of 1 is
	// used as the value for each count. If you include a Counts array, it must include
	// the same amount of values as the Values array.
	Counts []*float64
	// The value for the metric. Although the parameter accepts numbers of type Double,
	// CloudWatch rejects values that are either too small or too large. Values must be
	// in the range of -2^360 to 2^360. In addition, special values (for example, NaN,
	// +Infinity, -Infinity) are not supported.
	Value *float64
}

// This structure defines the metric to be returned, along with the statistics,
// period, and units.
type MetricStat struct {
	// The metric to return, including the metric name, namespace, and dimensions.
	Metric *Metric
	// The granularity, in seconds, of the returned data points. For metrics with
	// regular resolution, a period can be as short as one minute (60 seconds) and must
	// be a multiple of 60. For high-resolution metrics that are collected at intervals
	// of less than one minute, the period can be 1, 5, 10, 30, 60, or any multiple of
	// 60. High-resolution metrics are those metrics stored by a PutMetricData call
	// that includes a StorageResolution of 1 second. If the StartTime parameter
	// specifies a time stamp that is greater than 3 hours ago, you must specify the
	// period as follows or no data points in that time range is returned:
	//
	//     * Start
	// time between 3 hours and 15 days ago - Use a multiple of 60 seconds (1
	// minute).
	//
	//     * Start time between 15 and 63 days ago - Use a multiple of 300
	// seconds (5 minutes).
	//
	//     * Start time greater than 63 days ago - Use a multiple
	// of 3600 seconds (1 hour).
	Period *int32
	// When you are using a Put operation, this defines what unit you want to use when
	// storing the metric. In a Get operation, if you omit Unit then all data that was
	// collected with any unit is returned, along with the corresponding units that
	// were specified when the data was reported to CloudWatch. If you specify a unit,
	// the operation returns only data that was collected with that unit specified. If
	// you specify a unit that does not match the data collected, the results of the
	// operation are null. CloudWatch does not perform unit conversions.
	Unit StandardUnit
	// The statistic to return. It can include any CloudWatch statistic or extended
	// statistic.
	Stat *string
}

// This array is empty if the API operation was successful for all the rules
// specified in the request. If the operation could not process one of the rules,
// the following data is returned for each of those rules.
type PartialFailure struct {
	// The type of error.
	ExceptionType *string
	// The specified rule that could not be deleted.
	FailureResource *string
	// A description of the error.
	FailureDescription *string
	// The code of the error.
	FailureCode *string
}

// Specifies one range of days or times to exclude from use for training an anomaly
// detection model.
type Range struct {
	// The end time of the range to exclude. The format is yyyy-MM-dd'T'HH:mm:ss. For
	// example, 2019-07-01T23:59:59.
	EndTime *time.Time
	// The start time of the range to exclude. The format is yyyy-MM-dd'T'HH:mm:ss. For
	// example, 2019-07-01T23:59:59.
	StartTime *time.Time
}

// Represents a set of statistics that describes a specific metric.
type StatisticSet struct {
	// The minimum value of the sample set.
	Minimum *float64
	// The maximum value of the sample set.
	Maximum *float64
	// The number of samples used for the statistic set.
	SampleCount *float64
	// The sum of values for the sample set.
	Sum *float64
}

// A key-value pair associated with a CloudWatch resource.
type Tag struct {
	// A string that you can use to assign a value. The combination of tag keys and
	// values can help you organize and categorize your resources.
	Key *string
	// The value for the specified tag key.
	Value *string
}
